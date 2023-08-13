package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/systoms/wegin/backend/forms"
	"github.com/systoms/wegin/backend/services"
	"github.com/systoms/wegin/utils"
	"gorm.io/gorm"
	"net/http"
)

type Passport struct {
}

func (ctx *Passport) Login(c *gin.Context) {
	result := utils.Result{}
	var passportForm forms.LoginForm
	if err := c.ShouldBindJSON(&passportForm); err != nil {
		c.JSON(http.StatusOK, result.Error(500, err.Error()))
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	loginService := services.NewLoginService(db)
	token, err := loginService.Login(passportForm)
	if err != nil {
		c.JSON(http.StatusOK, result.Error(500, err.Error()))
	}

	c.JSON(http.StatusOK, result.Success(struct {
		Token string `json:"token"`
	}{
		Token: token,
	}))
}

// Logout 退出登录
func (ctx *Passport) Logout(c *gin.Context) {
	result := utils.Result{}
	token := c.GetHeader("Authorization")

	db := c.MustGet("db").(*gorm.DB)
	loginService := services.NewLoginService(db)
	err := loginService.Logout(token)

	if err != nil {
		c.JSON(http.StatusOK, result.Error(500, "退出失败"))
	}

	c.JSON(http.StatusOK, result.Success(nil))
}
