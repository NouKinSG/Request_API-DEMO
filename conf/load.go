package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

//如何配置映射成Config对象

// 从toml格式配置文件中加载配置
func LoadConfigFromToml(filePath string) error {
	config = NewDefaultConfig()
	// 读取Toml格式的配置
	_, err := toml.DecodeFile(filePath, config)
	if err != nil {
		return fmt.Errorf("load config from file error,path:%s,%s", filePath, err)
	}

	return nil
	// 第二种方法：return loadGloabal()
}

// 从环境变量中加载配置
func LoadConfigFromEnv() error {
	config = NewDefaultConfig()
	err := env.Parse(config)
	if err != nil {
		return err
	}

	return nil
	// 第二种方法：return loadGloabal()
}

// 加载全局实例
func loadGloabal() (err error) {
	// 加载db的全局实例
	db, err = config.MySQL.getDBConn()
	if err != nil {
		return
	}

	return
}
