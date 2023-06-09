package impl

import (
	"database/sql"

	"github.com/CMDB-Demo/apps/host"
	"github.com/CMDB-Demo/conf"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

//结构体是否满足我们的申明
//接口实现的静态检测
var _ host.Service = (*HostServiceImpl)(nil)

//构造logger函数
//logger 库抽象为接口了
//全局conf对象已经初始化 保证调用该函数之前。
func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{
		//host service 服务的子logger
		//使用封装的zap让其满足logger接口
		//为什么封装zap：
		//1.logger的全局实例
		//2.logger level的动态调整,logrus 不支持level动态调整
		//3.加入日志轮转功能的集合
		l:  zap.L().Named("Host"),
		db: conf.C().MySQL.GetDB(),
	}
}

//host的service的具体实现类
//所有根模块相关的功能都要放在这里面
type HostServiceImpl struct {
	l  logger.Logger
	db *sql.DB
}
