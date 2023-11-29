package impl

import (
	"context"

	"gitee.com/go-course/restful-api-demo-g7/apps/host"
)

// 完成对象核数据库直接转化
func (s *HostServiceImpl) save(ctx context.Context, ins *host.Host) error {
	// 1、校验数据合法性
	ins.Validate()

	return nil
}
