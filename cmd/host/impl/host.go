package impl

import (
	"context"

	"github.com/CMDB-Demo/cmd/host"
	"github.com/infraboard/mcube/logger"
)

//具体接口结构体的申明
func (i *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) (
	*host.Host, error) {
	//logger的用法
	//直接打印日志
	i.l.Named("Create").Error("create host")
	i.l.Error("create host")
	i.l.Info("create host")

	//带格式化日志打印
	i.l.Errorf("create host %s ", ins.Name)
	//如果需要带上一些metat信息。携带额外的meta数据常用与trace系统
	i.l.With(logger.NewAny("request-id", "req01")).Error("Create host with meta kv")
	return ins, nil
}

func (i *HostServiceImpl) QueryHost(ctx context.Context, req *host.QueryHostRequest) (
	*host.HostSet, error) {
	return nil, nil
}

func (i *HostServiceImpl) DescribeHost(ctx context.Context, req *host.DescribeHostRequest) (
	*host.Host, error) {
	return nil, nil
}

func (i *HostServiceImpl) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (
	*host.Host, error) {
	return nil, nil
}

func (i *HostServiceImpl) DeleteHost(ctx context.Context, req *host.DeleteHostRequest) (
	*host.Host, error) {
	return nil, nil
}
