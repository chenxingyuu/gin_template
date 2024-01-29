package datastore

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/chenxingyuu/gin_template/config"
	"github.com/chenxingyuu/gin_template/pkg/datastore/xmysql"
	"github.com/chenxingyuu/gin_template/pkg/datastore/xredis"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

// DS 抽象数据存储接口
type DS interface {
	init() error
}

type MySQLDS struct {
	Config *config.MySQLConfig
	DB     *xmysql.XDB
}

type RedisDS struct {
	Config *config.RedisConfig
	DB     *xredis.XDB
}

var mysqlOnce sync.Once
var redisOnce sync.Once

var MySQLClient = &MySQLDS{}
var RedisClient = &RedisDS{}

// Init MySQL客户端初始化
func (ds *MySQLDS) init() (err error) {
	if ds.DB == nil {
		mysqlOnce.Do(func() {
			var db *sql.DB
			// 创建数据库驱动连接
			db, _ = sql.Open("mysql", ds.Config.DSN())
			ds.DB = &xmysql.XDB{DB: db}
			if ds.DB != nil {
				ds.DB.SetMaxIdleConns(10)
				ds.DB.SetMaxOpenConns(100)
			}
		})
	}
	// 连接测试
	if err = ds.DB.Ping(); err != nil {
		return fmt.Errorf("failed to connect to MySQL: %v", err)
	}
	return
}

// Init Redis客户端初始化
func (ds *RedisDS) init() (err error) {
	if ds.DB == nil {
		redisOnce.Do(func() {
			var client *redis.Client
			client = redis.NewClient(&redis.Options{
				Addr:     ds.Config.Addr(),   // Redis服务器地址
				Password: ds.Config.Password, // Redis数据库密码，没有则留空
				DB:       ds.Config.Database, // 要连接的数据库编号
			})
			ds.DB = &xredis.XDB{Client: client}
		})
	}
	// 连接测试
	ctx := context.Background()
	if err = ds.DB.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("failed to connect to Redis: %v", err)
	}
	return
}

// InitDataStore 初始化数据存储客户端
func InitDataStore() {
	MySQLClient.Config = config.MySQL
	RedisClient.Config = config.Redis
	dss := []DS{MySQLClient, RedisClient}
	for index, ds := range dss {
		err := ds.init()
		if err != nil {
			panic(fmt.Errorf("failed to init datastore[%d]: %v", index, err))
		}
	}
	return
}
