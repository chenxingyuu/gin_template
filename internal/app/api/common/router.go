package common

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, apiPrefix string) {
	commonRoutes := r.Group(apiPrefix + "/common")
	{
		// ping
		commonRoutes.GET("/ping", PingHandler)
	}
}
