package main

import (
	"github.com/chenxingyuu/gin_template/config"
	"github.com/chenxingyuu/gin_template/internal/app/server"
	"github.com/chenxingyuu/gin_template/pkg/datastore"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 初始化配置
	config.InitConfig()

	// 初始化数据库
	datastore.InitDataStore()

	// 启动服务
	server.StartServer(config.Env.Addr)

	// 优雅退出
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 关闭服务
	server.ShutdownServer()
}
