package impl_test

import (
	"context"
	"fmt"
	"testing"

	"gitee.com/go-course/restful-api-demo-g7/apps/host"
	"gitee.com/go-course/restful-api-demo-g7/apps/host/impl"
	"gitee.com/go-course/restful-api-demo-g7/conf"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"
)

var (
	// 定义对象是满足该接口的实例
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

func init() {
	// 测试用例的配置文件
	err := conf.LoadConfigFromToml("../host/apps/etc/demo.toml")
	if err != nil {
		panic(err)
	}
	// 需要初始化Logger，
	//为什么不设计默认打印  -> 因为性能
	zap.DevelopmentSetup()

	// host service 的具体实现
	service = impl.NewHostServiceImpl()

}
