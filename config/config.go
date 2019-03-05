package config

import (
	"git.bitboolean.com/platformgo/misc/driver"
	"git.bitboolean.com/platformgo/misc/utils"
)

var (
	config GlobalConfig
)

type RedisConfig struct {
	Host string
	Port string
}

//GlobalConfig 全局配置
type GlobalConfig struct {
	HTTPBind string
	Mysql    driver.MysqlConfig
	TestUrl  string
	Redis   RedisConfig
}

// Config 返回配置文件
func Config() GlobalConfig {
	return config
}

// ParseConfig 解析配置文件
func ParseConfig(cfg string) {
	util.ParseConfig(cfg, &config)
}
