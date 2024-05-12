package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

var YamlConfig Yaml

func yamlConfigInit() {
	arr, err := readFile(" ", "application", "yaml")
	if err != nil {
		fmt.Println("读取配置文件失败")
		panic(err)
	}
	err = arr.Unmarshal(&YamlConfig)
	if err != nil {
		fmt.Println("转换配置文件失败")
		panic(err)
	}
}

/**
 * 通用的读取配置文件的方法，传入路径和文件名以及类型，返回一个Viper的指针
 */
func readFile(filePath, fileName, configType string) (config *viper.Viper, err error) {
	config = viper.New()
	config.AddConfigPath(filePath)
	config.SetConfigName(fileName)
	config.SetConfigType(configType)
	if err := config.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			panic("找不到配置文件")
		}
	}
	return
}
