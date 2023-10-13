package impl_test

import (
	"context"
	"testing"

	"gitee.com/go-course/restful-api-demo-g7/apps/host"
	"gitee.com/go-course/restful-api-demo-g7/apps/host/impl"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	// 定义对象是满足该接口的实例
	service host.Service
)

func TestCreate(t *testing.T) {
	ins := host.NewHost()
	ins.Name = "test"
	service.CreateHost(context.Background(), ins)
}

func init() {
	// 需要初始化Logger，
	//为什么不设计默认打印  -> 因为性能
	zap.DevelopmentSetup()

	// host service 的具体实现
	service = impl.NewHostServiceImpl()

}
