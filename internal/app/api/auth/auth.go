package auth

import (
	"github.com/chenxingyuu/gin_template/internal/app/service/auth"
	"github.com/chenxingyuu/gin_template/internal/app/service/user"
	"github.com/chenxingyuu/gin_template/pkg/response"
	"github.com/chenxingyuu/gin_template/pkg/xjwt"
	"github.com/gin-gonic/gin"
)

func PasswordLoginHandler(c *gin.Context) {
	// 参数校验
	var params PasswordLoginRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, response.StatusInvalidParams, err.Error())
		return
	}

	// 查询用户
	User, err := user.ByEmail(params.Email)
	if err != nil {
		response.Error(c, response.StatusUnauthorized, "Invalid credentials")
		return
	}

	// 验证密码
	if !auth.Authenticate(User, params.Password) {
		response.Error(c, response.StatusUnauthorized, "Invalid credentials")
		return
	}

	// 生成 token
	accessToken, err := xjwt.NewAccess(User)
	if err != nil {
		response.Error(c, response.StatusGenerateTokenError, "Failed to generate access token")
		return
	}

	// 生成 refresh token
	refreshToken, err := xjwt.NewRefresh(User)
	if err != nil {
		response.Error(c, response.StatusGenerateTokenError, "Failed to generate refresh token")
		return
	}

	response.Success(c, Token{
		Access:  accessToken,
		Refresh: refreshToken,
		Type:    "Bearer",
	})

	return
}
