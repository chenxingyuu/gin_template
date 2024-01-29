package common

import (
	"github.com/chenxingyuu/gin_template/pkg/response"
	"github.com/gin-gonic/gin"
)

// PingHandler 心跳检测
func PingHandler(c *gin.Context) {
	response.Success(c, nil)
}
