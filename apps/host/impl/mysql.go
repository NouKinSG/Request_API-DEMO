package impl

import (
	"gitee.com/go-course/restful-api-demo-g7/apps/host"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

// 接口实现的静态检查
var _ host.Service = (*HostServiceImpl)(nil)

func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{
		// Host service 服务的子logger
		// 封装的Zap让其满足 Logger接口
		// 为什么要封装
		// 		1、Logger全局实例
		// 		2、Logger Level的动态调整，Logrus不支持Level共同调整
		// 		3、加入日志轮转功能的集合
		l: zap.L().Named("Host"),
	}
}

type HostServiceImpl struct {
	l logger.Logger
}
