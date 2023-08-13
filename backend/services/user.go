package services

import (
	"github.com/systoms/wegin/backend/models"
	"github.com/systoms/wegin/utils"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// GenerateToken 生成用户token
func (ctx *UserService) GenerateAccessToken(model models.UserModel) models.UserModel {
	token := utils.GenerateRandomString(16)
	ctx.db.Model(model).Update("access_token", token)
	return model
}

// GetOneByUsername 根据用户名查询用户
func (ctx *UserService) GetOneByUsername(username string) models.UserModel {
	var userModel models.UserModel
	ctx.db.First(&userModel, "username = ?", username)
	return userModel
}

// GetOneByAccessToken 根据access_token查询用户
func (ctx *UserService) GetOneByAccessToken(access_token string) models.UserModel {
	var userModel models.UserModel
	ctx.db.First(&userModel, "access_token = ?", access_token)
	return userModel
}

// ValidatedPassword 验证密码
func (ctx *UserService) ValidatedPassword(userModel models.UserModel, password string) bool {
	return utils.CalculateMD5(password) != userModel.Password
}

// ResetAccessToken 重置access_token
func (ctx *UserService) ResetAccessToken(userModel models.UserModel) bool {
	ctx.db.Model(userModel).Update("access_token", "")
	return true
}
