package impl

import (
	"database/sql"

	"gitee.com/go-course/restful-api-demo-g7/apps/host"
	"gitee.com/go-course/restful-api-demo-g7/conf"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

// 接口实现的静态检查
var _ host.Service = (*HostServiceImpl)(nil)

// NewHostServiceImpl    保证调用函数之前，全局的Logger已经被初始化
func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{
		// Host service 服务的子logger
		// 封装的Zap让其满足 Logger接口
		// 为什么要封装
		// 		1、Logger全局实例
		// 		2、Logger Level的动态调整，Logrus不支持Level共同调整
		// 		3、加入日志轮转功能的集合
		l:  zap.L().Named("Host"),
		db: conf.C().MySQL.GetDB(),
	}
}

type HostServiceImpl struct {
	l  logger.Logger
	db *sql.DB
}
