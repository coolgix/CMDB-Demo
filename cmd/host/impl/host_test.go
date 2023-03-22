package impl_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/CMDB-Demo/cmd/host"
	"github.com/CMDB-Demo/cmd/host/impl"
	"github.com/infraboard/mcube/logger/zap"
)

//全局var

var (
	//定义这个对象是满足该接口的实例
	service host.Service
)

func TestCreate(t *testing.T) {
	ins := host.NewHost()
	ins.Name = "test"
	service.CreateHost(context.Background(), ins)
}

//初始化测试用例
func init() {
	//没有日志需要初始化全局logger
	// 为什么不设置为默认打印，因为性能
	fmt.Println(zap.DevelopmentSetup())

	//接口的具体实现
	//host service
	service = impl.NewHostServiceImpl()

}
