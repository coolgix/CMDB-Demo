package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

//如何把配置文件映射为config对象

//从配置文件加载
func LoadConfigFromToml(filePath string) error {
	config = NewDefaultConfig()

	//根据工具的用法需要先读取toml格式的配置文件
	_, err := toml.DecodeFile(filePath, config)
	if err != nil {
		return fmt.Errorf("load config from file error, path:%s, %s", filePath, err)
	}

	return nil
}

//从环境变量加载
func LoadConfigFromEnv() error {
	config = NewDefaultConfig()

	err := env.Parse(config)
	if err != nil {
		return err
	}

	return nil
}
