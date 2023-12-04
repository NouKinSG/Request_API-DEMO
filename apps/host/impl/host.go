package impl

import (
	"context"

	"gitee.com/go-course/restful-api-demo-g7/apps/host"
	"github.com/infraboard/mcube/logger"
)

// 业务处理层（Controller层）的实现
func (i *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	// 直接打印日志
	i.l.Debug("create host")
	// 带Format的日志打印，fmt.Sprintf()
	i.l.Debugf("create host %s", ins.Name)
	// 携带额外mata数据，常用于Trace系统
	i.l.With(logger.NewAny("request-id", "req01")).Debug("create host with meta kv")

	// 1、校验数据合法性
	if err := ins.Validate(); err != nil {
		return nil, err
	}

	// 2、处理默认值
	ins.InjectDefault()

	// 有dao模块  负责  把对象入库
	if err := i.save(ctx, ins); err != nil {
		return nil, err
	}
	return ins, nil
}

func (i *HostServiceImpl) QueryHost(ctx context.Context, req *host.QueryHostRequest) (*host.HostSet, error) {
	return nil, nil
}

func (i *HostServiceImpl) DescribeHost(ctx context.Context, req *host.QueryHostRequest) (*host.Host, error) {
	return nil, nil
}

func (i *HostServiceImpl) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (*host.Host, error) {
	return nil, nil
}

func (i *HostServiceImpl) DeleteHost(ctx context.Context, req *host.DeleteHostRequest) (*host.Host, error) {
	return nil, nil
}
