package datastore

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/chenxingyuu/gin_template/config"
	"github.com/chenxingyuu/gin_template/pkg/datastore/xmysql"
	"github.com/chenxingyuu/gin_template/pkg/datastore/xredis"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitMySQLClient(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	assert.NoError(t, err)

	sqlMock.ExpectPing().WillReturnError(nil)

	mockClient := &MySQLDS{DB: &xmysql.XDB{DB: dbMock}}
	initErr := mockClient.init()
	assert.NoError(t, initErr)
}

func TestInitMySQLClientWithPingError(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	assert.NoError(t, err)

	sqlMock.ExpectPing().WillReturnError(fmt.Errorf("ping error"))

	mockClient := &MySQLDS{DB: &xmysql.XDB{DB: dbMock}}
	initErr := mockClient.init()
	assert.Error(t, initErr)
	assert.EqualError(t, initErr, "failed to connect to MySQL: ping error")
}

func TestInitRedisClient(t *testing.T) {
	dbMock, mock := redismock.NewClientMock()
	mock.ExpectPing().SetVal("PONG")

	mockClient := &RedisDS{DB: &xredis.XDB{Client: dbMock}}
	initErr := mockClient.init()
	assert.NoError(t, initErr)
}

func TestInitRedisClientWithPingError(t *testing.T) {
	dbMock, mock := redismock.NewClientMock()
	mock.ExpectPing().SetErr(fmt.Errorf("ping error"))

	mockClient := &RedisDS{DB: &xredis.XDB{Client: dbMock}}
	initErr := mockClient.init()
	assert.Error(t, initErr)
	assert.EqualError(t, initErr, "failed to connect to Redis: ping error")
}

func TestInitDataStore(t *testing.T) {
	// 加载配置文件
	config.Env.ConfigPath = "../../config.yaml"
	config.InitConfig()
	InitDataStore()
}
