package xjwt

import (
	"errors"
	"github.com/chenxingyuu/gin_template/config"
	"github.com/chenxingyuu/gin_template/internal/app/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenType uint8

const (
	Refresh TokenType = iota
	Access
)

type Claims struct {
	UserID    int64     `json:"user_id"`
	Username  string    `json:"username"`
	TokenType TokenType `json:"token_type"`
	jwt.StandardClaims
}

// New 生成 JWT
func New(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(config.Jwt.GetSecretKey())
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func NewAccess(user *model.User) (string, error) {
	return New(Claims{
		UserID:    int64(user.ID),
		Username:  user.Username,
		TokenType: Access,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(config.Jwt.GetAccessExp()).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})
}

func NewRefresh(user *model.User) (string, error) {
	return New(Claims{
		UserID:    int64(user.ID),
		Username:  user.Username,
		TokenType: Refresh,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(config.Jwt.GetRefreshExp()).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})
}

func Validate(t string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(t, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return config.Jwt.GetSecretKey(), nil
	})

	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token已过期")
			}
		}
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("无效的Token")
}

func ValidateAccess(t string) (*Claims, error) {
	claims, err := Validate(t)
	if err == nil && claims.TokenType == Access {
		return claims, nil
	}
	return nil, errors.New("token类型错误")
}

func ValidateRefresh(t string) (*Claims, error) {
	claims, err := Validate(t)
	if err == nil && claims.TokenType == Refresh {
		return claims, nil
	}
	return nil, errors.New("token类型错误")
}
