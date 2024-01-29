package config

import (
	"os"
	"path/filepath"
	"testing"
)

var validConfigStr = `
mysql:
  host: 127.0.0.1
  port: 3306
  username: root
  password: 123456
  database: hello
  charset: utf8mb4

redis:
  host: 127.0.0.1
  port: 6379
  password: 123456
  database: 1

jwt:
  secretKey: hello
  accessExp: 3600
  refreshExp: 7200
`

func TestLoadConfig(t *testing.T) {
	t.Run("LoadConfigWithValidPath", func(t *testing.T) {
		// 创建临时配置文件
		configPath := filepath.Join(os.TempDir(), "config.yaml")
		_ = os.WriteFile(configPath, []byte(validConfigStr), 0644)

		loadConfig(configPath)
		if MySQL == nil || Redis == nil || Jwt == nil {
			t.Errorf("Expected MySQL, Redis, and Jwt to be not nil")
		}
	})

	t.Run("LoadConfigWithInvalidPath", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected a panic for invalid config path")
			}
		}()
		loadConfig("invalid/path/to/config.yaml")
	})
}

func TestInitConfig(t *testing.T) {
	t.Run("InitConfigWithValidPath", func(t *testing.T) {
		// 创建临时配置文件
		configPath := filepath.Join(os.TempDir(), "config.yaml")
		_ = os.WriteFile(configPath, []byte(validConfigStr), 0644)

		InitConfig(configPath)
		if MySQL == nil || Redis == nil || Jwt == nil {
			t.Errorf("Expected MySQL, Redis, and Jwt to be not nil")
		}
	})

	t.Run("InitConfigWithInvalidPath", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected a panic for invalid config path")
			}
		}()
		InitConfig("invalid/path/to/config.yaml")
	})
}
