package conf

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//定义我们程序需要的配置对象

//定义包的全局config实例对象
//也就是我们程序在内存中的配置对象
// Config 应用配置
//封装为一个对象，来与外部进行对接
//通过标签来与toml文件一一对应
//程序内部获取配置，都通过读取该对象
//该对象 什么时候初始化？
//配置加载的是：
// LoadConfigFromToml
// LoadConfigFromEnv
//全局的默认Config对象
//为了程序在运行的时候不被恶意修改，设置为私有变量
var config *Config

// 全局MySQL 客户端实例
var db *sql.DB

//如过想获得配置，我们单独提供一个函数
//全局Config 对象获取函数
func C() *Config {
	return config
}

//初始化一个默认的config对象
func NewDefaultConfig() *Config {
	return &Config{
		App:   NewDefaultApp(),
		Log:   NewDefaultLog(),
		MySQL: NewDefaultMySQL(),
	}
}

// Config 应用配置
// 通过封装为一个对象, 来与外部配置进行对接
type Config struct {
	App   *App   `toml:"app"`
	Log   *Log   `toml:"log"`
	MySQL *MySQL `toml:"mysql"`
}

//mysql的默认值
func NewDefaultApp() *App {
	return &App{
		Name: "demo",
		Host: "127.0.0.1",
		Port: "8080",
	}
}

type App struct {
	Name string `toml:"name" env:"APP_NAME"`
	Host string `toml:"host" env:"APP_HOST"`
	Port string `toml:"port" env:"APP_PORT"`
	Key  string `toml:"key" env:"APP_KEY"`
	//ssl配置暂时不用

	//EnableSSL bool   `toml:"enable_ssl" env:"APP_ENABLE_SSL"`
	//CertFile  string `toml:"cert_file" env:"APP_CERT_FILE"`
	//KeyFile   string `toml:"key_file" env:"APP_KEY_FILE"`
}

func NewDefaultMySQL() *MySQL {
	return &MySQL{
		Host:     "127.0.0.1",
		Port:     "3306",
		UserName: "demo",
		Password: "123456",
		Database: "demo",
		//生产环境配置
		MaxOpenConn: 200,
		MaxIdleConn: 100,
	}
}

type MySQL struct {
	//为了配合读取toml和enc 添加了相关的标签
	Host     string `toml:"host" env:"MYSQL_HOST"`
	Port     string `toml:"port" env:"MYSQL_PORT"`
	UserName string `toml:"username" env:"MYSQL_USERNAME"`
	Password string `toml:"password" env:"MYSQL_PASSWORD"`
	Database string `toml:"database" env:"MYSQL_DATABASE"`
	//因为使用mysql的连接池需要对一些池做一些规划配置
	//跟mysql服务端相关，控制当前程序的mysql的打开的连接数
	MaxOpenConn int `toml:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	//控制mysql连接的复用,最多允许5个复用
	MaxIdleConn int `toml:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	//链接的生命周期，这个不能大于mysql的生命周期
	//网络问题抖动，一个链接12h ，换一个connect，验证可用性
	MaxLifeTime int `toml:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	//最大的一个复用链接时间过了时间会直接回收
	MaxIdleTime int `toml:"max_idle_time" env:"MYSQL_MAX_idle_TIME"`

	// 作为私有变量, 用户与控制GetDB
	lock sync.Mutex
}

//分装一个链接db的全局实例
// 1. 第一种方式, 使用LoadGlobal 在加载时 初始化全局db实例
// 2. 第二种方式, 惰性加载, 获取DB是，动态判断再初始化
//防止多个进行竞争需要加锁处理
func (m *MySQL) GetDB() *sql.DB {
	// 直接加锁, 锁住临界区
	m.lock.Lock()
	defer m.lock.Unlock()

	// 如果实例不存在, 会比较意外，会报错，就初始化一个新的实例
	if db == nil {
		conn, err := m.getDBConn()
		if err != nil {
			panic(err)
		}
		db = conn
	}

	// 全局变量db就一定存在了
	return db
}

//抽象一个mysql的客户端
// 连接池, driverConn具体的连接对象, 他维护着一个Socket
// pool []*driverConn, 维护pool里面的连接都是可用的, 定期检查我们的conn健康情况
// 某一个driverConn已经失效, driverConn.Reset(), 清空该结构体的数据, Reconn获取一个连接, 让该conn借壳存活
// 避免driverConn结构体的内存申请和释放的一个成本
func (m *MySQL) getDBConn() (*sql.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&multiStatements=true", m.UserName, m.Password, m.Host, m.Port, m.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect to mysql<%s> error, %s", dsn, err.Error())
	}

	db.SetMaxOpenConns(m.MaxOpenConn)
	db.SetMaxIdleConns(m.MaxIdleConn)
	db.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	db.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql<%s> error, %s", dsn, err.Error())
	}
	return db, nil
}

//默认log配置
func NewDefaultLog() *Log {
	return &Log{
		//debug info error warn
		Level:  "info",
		Format: TextFormat,
		To:     ToStdout,
	}
}

//用于配置全局的log对象
type Log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"LOG_TO"`
	PathDir string    `toml:"path_dir" env:"LOG_PATH_DIR"`
}
