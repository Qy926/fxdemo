// lib/config/load.go
package config

import (
	"fmt"
	"github.com/spf13/viper"
)

//文件使用Viper框架从YML文件加载配置
func getViper() *viper.Viper {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigFile("config.yml")
	return v
}

func NewConfig() (*Config, error) {
	fmt.Println("Loading configuration")
	v := getViper()
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var config Config
	err = v.Unmarshal(&config)
	return &config, err
}
