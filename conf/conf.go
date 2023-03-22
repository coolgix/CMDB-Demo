package conf

//定义我们程序需要的配置对象

// Config 应用配置
//封装为一个对象，来与外部进行对接
//通过标签来与toml文件一一对应
type Config struct {
	App   *app   `toml:"app"`
	Log   *Log   `toml:"log"`
	Mysql *Mysql `toml:"mysql"`
}

type app struct {
	Name string `toml:"name" env:"APP_NAME"`
	Host string `toml:"host" env:"APP_HOST"`
	Port string `toml:"port" env:"APP_PORT"`
	Key  string `toml:"key" env:"APP_KEY"`
	//ssl配置暂时不用

	//EnableSSL bool   `toml:"enable_ssl" env:"APP_ENABLE_SSL"`
	//CertFile  string `toml:"cert_file" env:"APP_CERT_FILE"`
	//KeyFile   string `toml:"key_file" env:"APP_KEY_FILE"`
}
type Mysql struct {
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
	//lock        sync.Mutex
}

type Log struct {
	Level   string `toml:"level" env:"LOG_LEVEL"`
	PathDir string `toml:"path_dir" env:"LOG_PATH_DIR"`
	//Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	//To LogTo `toml:"to" env:"LOG_TO"`
}
