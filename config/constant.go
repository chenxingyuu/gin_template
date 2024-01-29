package config

import (
	"fmt"
	"time"
)

type EnvConfig struct {
	Addr       string `env:"SERVER_ADDR" default:"127.0.0.1:8080" desc:"server address"`
	ProjectEnv string `env:"PROJECT_ENV" default:"dev" desc:"project environment"`
	APIVersion string `env:"API_VERSION" default:"v1" desc:"api version"`
	ConfigPath string `env:"CONFIG_PATH" default:"config.yaml" desc:"config file path"`

	WriteTimeout      int `env:"WRITE_TIMEOUT" default:"30" desc:"write timeout"`
	ReadTimeout       int `env:"READ_TIMEOUT" default:"30" desc:"read timeout"`
	ReadHeaderTimeout int `env:"READ_HEADER_TIMEOUT" default:"30" desc:"read header timeout"`
}

type MySQLConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Charset  string `json:"charset"`
}

func (c MySQLConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", c.Username, c.Password, c.Host, c.Port, c.Database, c.Charset)
}

type RedisConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Database int    `json:"database"`
}

func (c RedisConfig) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

type JwtConfig struct {
	SecretKey  string        `json:"secret_key"`
	AccessExp  time.Duration `json:"access_exp"`
	RefreshExp time.Duration `json:"refresh_exp"`
}

func (c JwtConfig) GetSecretKey() []byte {
	return []byte(c.SecretKey)
}

func (c JwtConfig) GetAccessExp() time.Duration {
	return time.Second * c.AccessExp
}

func (c JwtConfig) GetRefreshExp() time.Duration {
	return time.Second * c.RefreshExp
}
