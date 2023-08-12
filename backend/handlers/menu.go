package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/systoms/wegin/utils"
	"net/http"
)

type Menu struct {
}

func (ctx *Menu) List(c *gin.Context) {
	result := utils.Result{}
	c.JSON(http.StatusOK, result.Success(nil))
}

func (ctx *Menu) Create(c *gin.Context) {
	result := utils.Result{}
	c.JSON(http.StatusOK, result.Success(nil))
}

func (ctx *Menu) Update(c *gin.Context) {
	result := utils.Result{}
	c.JSON(http.StatusOK, result.Success(nil))
}

func (ctx *Menu) Info(c *gin.Context) {
	result := utils.Result{}
	c.JSON(http.StatusOK, result.Success(nil))
}

func (ctx *Menu) Delete(c *gin.Context) {
	result := utils.Result{}
	c.JSON(http.StatusOK, result.Success(nil))
}

func (ctx *Menu) Tree(c *gin.Context) {
	result := utils.Result{}
	c.JSON(http.StatusOK, result.Success(nil))
}
