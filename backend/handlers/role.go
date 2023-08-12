package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/systoms/wegin/utils"
	"net/http"
)

type Role struct {
}

func (ctx *Role) List(c *gin.Context) {
	result := utils.Result{}
	c.JSON(http.StatusOK, result.Success(nil))
}

func (ctx *Role) Create(c *gin.Context) {
	result := utils.Result{}
	c.JSON(http.StatusOK, result.Success(nil))
}

func (ctx *Role) Update(c *gin.Context) {
	result := utils.Result{}
	c.JSON(http.StatusOK, result.Success(nil))
}

func (ctx *Role) Info(c *gin.Context) {
	result := utils.Result{}
	c.JSON(http.StatusOK, result.Success(nil))
}

func (ctx *Role) Delete(c *gin.Context) {
	result := utils.Result{}
	c.JSON(http.StatusOK, result.Success(nil))
}
