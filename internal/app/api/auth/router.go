package auth

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, apiPrefix string) {
	authRoutes := r.Group(apiPrefix + "/auth")
	{
		// auth login
		authRoutes.POST("/login", PasswordLoginHandler)
	}
}
