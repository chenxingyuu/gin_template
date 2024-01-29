package config

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
)

var (
	MySQL *MySQLConfig
	Redis *RedisConfig
	Jwt   *JwtConfig
)

// loadConfig 读取配置文件
func loadConfig(configPath string) {
	v := viper.New()

	// 设置配置文件路径
	v.SetConfigFile(configPath)
	// 读取配置文件
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// 读取MySQL配置
	if err := v.UnmarshalKey("mysql", &MySQL); err != nil {
		panic(fmt.Errorf("failed to unmarshal mysql config: %s", err))
	}
	// 读取Redis配置
	if err := v.UnmarshalKey("redis", &Redis); err != nil {
		panic(fmt.Errorf("failed to unmarshal redis config: %s", err))
	}
	// 读取Jwt配置
	if err := v.UnmarshalKey("jwt", &Jwt); err != nil {
		panic(fmt.Errorf("failed to unmarshal jwt config: %s", err))
	}
}

// InitConfig 初始化配置文件
func InitConfig(path ...string) {
	var configPath string
	// 获取配置文件路径
	if len(path) > 0 {
		configPath = path[0]
	} else {
		// 默认配置文件路径
		configPath = Env.ConfigPath
	}
	// 获取绝对路径
	configPath, _ = filepath.Abs(configPath)
	// 读取配置文件
	loadConfig(configPath)
}
