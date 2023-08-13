package services

import (
	"errors"
	"github.com/systoms/wegin/backend/forms"
	"gorm.io/gorm"
)

type LoginService struct {
	db *gorm.DB
}

func NewLoginService(db *gorm.DB) *LoginService {
	return &LoginService{db: db}
}

// Login 实现用户登录操作
func (ctx LoginService) Login(loginForm forms.LoginForm) (string, error) {
	userService := NewUserService(ctx.db)

	userModel := userService.GetOneByUsername(loginForm.Username)
	if userModel.ID <= 0 {
		return "", errors.New("用户不存在")
	}
	if !userService.ValidatedPassword(userModel, userModel.Password) {
		return "", errors.New("密码错误")
	}

	userService.GenerateAccessToken(userModel)

	return userModel.AccessToken, nil
}

// Logout 实现用户注销操作
func (ctx LoginService) Logout(accessToken string) error {
	userService := NewUserService(ctx.db)

	userModel := userService.GetOneByAccessToken(accessToken)
	if userModel.ID <= 0 {
		return errors.New("用户不存在")
	}

	userService.ResetAccessToken(userModel)
	return nil
}
