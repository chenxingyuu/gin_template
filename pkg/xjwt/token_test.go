package xjwt

import (
	"github.com/chenxingyuu/gin_template/config"
	"github.com/chenxingyuu/gin_template/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var mockUser = &model.User{
	ID:       1,
	Username: "chenxingyu",
}

func TestNewAccess(t *testing.T) {
	config.InitConfig()

	token, newErr := NewAccess(mockUser)
	assert.NoError(t, newErr)

	claims, valErr := Validate(token)
	assert.NoError(t, valErr)
	assert.Equal(t, int64(mockUser.ID), claims.UserID)
	assert.Equal(t, mockUser.Username, claims.Username)
	assert.Equal(t, Access, claims.TokenType)
}

func TestNewRefresh(t *testing.T) {
	config.InitConfig()

	token, newErr := NewRefresh(mockUser)
	assert.NoError(t, newErr)

	claims, valErr := Validate(token)
	assert.NoError(t, valErr)
	assert.Equal(t, int64(mockUser.ID), claims.UserID)
	assert.Equal(t, mockUser.Username, claims.Username)
	assert.Equal(t, Refresh, claims.TokenType)
}

func TestValidate(t *testing.T) {
	config.InitConfig()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2ODY4ODM2MzcsImlhdCI6MTY4Njg4MzYxM30.Um6OmpknnJLpnny8jAdgZiTFWcx9OxxaRO5A-Qn-_6o"

	_, err := Validate(token)
	assert.Error(t, err)
	assert.EqualError(t, err, "token已过期")

	errToken := "12345"
	_, err = Validate(errToken)
	assert.Error(t, err)
	assert.EqualError(t, err, "token contains an invalid number of segments")
}

func TestValidateAccess(t *testing.T) {
	config.InitConfig()

	token, _ := NewAccess(mockUser)

	_, err := ValidateAccess(token)
	assert.NoError(t, err)

	errToken, _ := NewRefresh(mockUser)
	_, err = ValidateAccess(errToken)
	assert.Error(t, err)
	assert.EqualError(t, err, "token类型错误")
}

func TestValidateRefresh(t *testing.T) {
	config.InitConfig()

	token, _ := NewRefresh(mockUser)

	_, err := ValidateRefresh(token)
	assert.NoError(t, err)

	errToken, _ := NewAccess(mockUser)
	_, err = ValidateRefresh(errToken)
	assert.Error(t, err)
	assert.EqualError(t, err, "token类型错误")
}
