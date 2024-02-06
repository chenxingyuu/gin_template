package auth

import (
	"github.com/chenxingyuu/gin_template/internal/app/model"
	"golang.org/x/crypto/bcrypt"
)

// Authenticate 验证用户
func Authenticate(user *model.User, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil
}

// EncryptPassword 对密码进行加密
func EncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}
