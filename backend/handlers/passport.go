package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/systoms/wegin/utils"
	"net/http"
)

type Passport struct {
}

func (ctx *Passport) Login(c *gin.Context) {
	result := utils.Result{}
	c.JSON(http.StatusOK, result.Success(nil))
}

func (ctx *Passport) Logout(c *gin.Context) {
	result := utils.Result{}
	c.JSON(http.StatusOK, result.Success(nil))
}
