package user

import "github.com/chenxingyuu/gin_template/internal/app/model"

func ByEmail(email string) (user *model.User, err error) {
	user = &model.User{
		ID:       1,
		Username: "chenxingyu",
		Email:    "chenxingyu@ndnu.edu.cn",
		Password: "$2a$10$rarHg7dD3Rx0SLFxVLUpBORLP7a1hAqWQmAKmpDwHUcBfuQqbR386",
	}
	return
}
