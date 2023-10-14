package conf_test

import (
	"fmt"
	"os"
	"testing"

	"gitee.com/go-course/restful-api-demo-g7/conf"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfigFromToml(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/demo.toml")
	if should.NoError(err) {
		should.Equal(conf.C().App.Name, "demo")
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	should := assert.New(t)
	os.Setenv("MYSQL_DATABASE", "unit_test")

	err := conf.LoadConfigFromEnv()
	if should.NoError(err) {
		should.Equal(conf.C().MySQL.Database, "unit_test")
		fmt.Println(conf.C().MySQL.Database)
	}
}
