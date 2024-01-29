package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type jsonOption struct {
	Code    int
	Message string
	Data    interface{}
}

func getMessage(code int, messages ...string) string {
	messageStr := Message(code)
	if messages != nil {
		messageStr = messageStr + ": " + strings.Join(messages, ", ")
	}
	return messageStr
}

func json(c *gin.Context, statusCode int, option jsonOption) {
	c.JSON(statusCode, gin.H{
		"code":    option.Code,
		"message": option.Message,
		"data":    option.Data,
	})
}

func Error(c *gin.Context, code int, messages ...string) {
	messageStr := getMessage(code, messages...)
	json(c, http.StatusBadRequest, jsonOption{
		Code:    code,
		Message: messageStr,
	})
}

func Success(c *gin.Context, data interface{}) {
	json(c, http.StatusOK, jsonOption{
		Code:    StatusOK,
		Message: Message(StatusOK),
		Data:    data,
	})
}

func Unauthorized(c *gin.Context, code int, messages ...string) {
	messageStr := getMessage(code, messages...)
	json(c, http.StatusUnauthorized, jsonOption{
		Code:    code,
		Message: messageStr,
	})
}
