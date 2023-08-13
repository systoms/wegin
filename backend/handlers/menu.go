package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/systoms/wegin/backend/forms"
	"github.com/systoms/wegin/backend/services"
	"github.com/systoms/wegin/utils"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Menu struct {
}

func (ctx *Menu) List(c *gin.Context) {
	result := utils.Result{}

	db := c.MustGet("db").(*gorm.DB)
	menuService := services.NewMenuService(db)
	page := c.PostForm("page")

	newPage, err := strconv.Atoi(page)
	if err != nil {
		newPage = 1
	}
	data, err := menuService.GetList(newPage)
	if err != nil {
		c.JSON(http.StatusOK, result.Error(500, err.Error()))
	}

	c.JSON(http.StatusOK, result.Success(data))
}

func (ctx *Menu) Create(c *gin.Context) {
	result := utils.Result{}
	var menuCreateForm forms.MenuCreateForm
	if err := c.ShouldBindJSON(&menuCreateForm); err != nil {
		c.JSON(http.StatusOK, result.Error(500, err.Error()))
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	menuService := services.NewMenuService(db)
	menuModel, err := menuService.Create(menuCreateForm)
	if err != nil {
		c.JSON(http.StatusOK, result.Error(500, err.Error()))
	}

	c.JSON(http.StatusOK, result.Success(menuModel))
}

func (ctx *Menu) Update(c *gin.Context) {
	result := utils.Result{}
	var menuUpdateForm forms.MenuUpdateForm
	if err := c.ShouldBindJSON(&menuUpdateForm); err != nil {
		c.JSON(http.StatusOK, result.Error(500, err.Error()))
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	menuService := services.NewMenuService(db)
	menuModel, err := menuService.Update(menuUpdateForm)
	if err != nil {
		c.JSON(http.StatusOK, result.Error(500, err.Error()))
	}

	c.JSON(http.StatusOK, result.Success(menuModel))
}

func (ctx *Menu) Info(c *gin.Context) {
	result := utils.Result{}
	db := c.MustGet("db").(*gorm.DB)
	menuService := services.NewMenuService(db)
	id := c.PostForm("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, result.Error(500, "请选择菜单"))
	}
	data, err := menuService.GetOneById(newId)
	if err != nil {
		c.JSON(http.StatusOK, result.Error(500, err.Error()))
	}

	c.JSON(http.StatusOK, result.Success(data))
}

func (ctx *Menu) Delete(c *gin.Context) {
	result := utils.Result{}
	db := c.MustGet("db").(*gorm.DB)
	menuService := services.NewMenuService(db)
	id := c.PostForm("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, result.Error(500, "请选择菜单"))
	}
	menuModel, err := menuService.GetOneById(newId)
	if err != nil {
		c.JSON(http.StatusOK, result.Error(500, err.Error()))
	}

	err = menuService.Delete(menuModel)
	if err != nil {
		c.JSON(http.StatusOK, result.Error(500, err.Error()))
	}

	c.JSON(http.StatusOK, result.Success(nil))
}

func (ctx *Menu) Tree(c *gin.Context) {
	result := utils.Result{}
	db := c.MustGet("db").(*gorm.DB)
	menuService := services.NewMenuService(db)

	menuNodes, err := menuService.Tree()
	if err != nil {
		c.JSON(http.StatusOK, result.Error(500, err.Error()))
	}

	c.JSON(http.StatusOK, result.Success(menuNodes))
}
