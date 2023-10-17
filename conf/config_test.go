package conf_test

import (
	"os"
	"testing"

	"gitee.com/go-course/restful-api-demo-g7/conf"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfigFromToml(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/demo.toml")
	if should.NoError(err) {
		should.Equal("root", conf.C().App.Name)
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	should := assert.New(t)
	os.Setenv("MYSQL_DATABASE", "unit_test")

	err := conf.LoadConfigFromEnv()
	if should.NoError(err) {
		should.Equal("unit_test", conf.C().MySQL.Database)
		// fmt.Println(conf.C().MySQL.Database)
	}
}

func TestGetDB(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/demo.toml")
	if should.NoError(err) {

		conf.C().MySQL.GetDB()
	}
	// 通过环境变量读取
	// os.Setenv("MYSQL_DATABASE", "unit_test")

	// err := conf.LoadConfigFromEnv()
	// if should.NoError(err) {
	// 	conf.C().MySQL.GetDB()
	// }
}
