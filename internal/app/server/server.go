package server

import (
	"context"
	"errors"
	"github.com/chenxingyuu/gin_template/config"
	"github.com/chenxingyuu/gin_template/internal/app/api/common"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

const (
	maxWaitTimeBeforeShutdown = 10
)

var srv *http.Server

// StartServer 启动服务
func StartServer(addr string) {
	// 设置 gin 模式
	setGinMode()

	// 创建 gin 实例
	app := gin.Default()

	// 设置路由
	setupRouter(app)

	// 创建 http server
	srv = &http.Server{
		Addr:              addr,
		Handler:           app,
		ReadTimeout:       time.Duration(config.Env.ReadTimeout) * time.Second,
		ReadHeaderTimeout: time.Duration(config.Env.ReadHeaderTimeout) * time.Second,
		WriteTimeout:      time.Duration(config.Env.WriteTimeout) * time.Second,
	}

	// 启动服务
	go func() {
		log.Println("Start Http Server ", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Panicln("Failed to serve: ", err)
		}
	}()
}

// ShutdownServer 关闭服务
func ShutdownServer() {
	// 等待请求结束
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*maxWaitTimeBeforeShutdown)
	// 释放资源
	defer cancel()
	// 关闭服务
	if err := srv.Shutdown(ctx); err != nil {
		log.Print("shutdown server: ", err)
	}
}

// setGinMode 设置 gin 模式
func setGinMode() {
	env := config.Env.ProjectEnv
	if env == "dev" {
		gin.SetMode(gin.DebugMode)
	} else if env == "test" {
		gin.SetMode(gin.TestMode)
	} else if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
}

// setupRouter 设置路由
func setupRouter(app *gin.Engine) {
	// 设置 api 前缀
	apiPrefix := "/api/" + config.Env.APIVersion
	// 设置路由
	common.SetupRoutes(app, apiPrefix)
}
