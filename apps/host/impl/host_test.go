package impl_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/CMDB-Demo/apps/host"
	"github.com/CMDB-Demo/apps/host/impl"
	"github.com/CMDB-Demo/conf"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"
)

//全局var

var (
	//定义这个对象是满足该接口的实例
	service host.Service
)

func TestCreate(t *testing.T) {
	should := assert.New(t)

	ins := host.NewHost()
	ins.Name = "test"
	ins, err := service.CreateHost(context.Background(), ins)
	if should.NoError(err) {
		fmt.Println(ins)
	}
}

//初始化测试用例
func init() {
	//测试用例的配置文件
	err := conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}
	//没有日志需要初始化全局logger
	// 为什么不设置为默认打印，因为性能
	fmt.Println(zap.DevelopmentSetup())

	//接口的具体实现
	//host service
	service = impl.NewHostServiceImpl()

}
