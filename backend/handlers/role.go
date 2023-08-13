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

type Role struct {
}

func (ctx *Role) List(c *gin.Context) {
	result := utils.Result{}

	db := c.MustGet("db").(*gorm.DB)
	roleService := services.NewRoleService(db)
	page := c.PostForm("page")

	newPage, err := strconv.Atoi(page)
	if err != nil {
		newPage = 1
	}
	data, err := roleService.GetList(newPage)
	if err != nil {
		c.JSON(http.StatusOK, result.Error(500, err.Error()))
	}

	c.JSON(http.StatusOK, result.Success(data))
}

func (ctx *Role) Create(c *gin.Context) {
	result := utils.Result{}
	var roleCreateForm forms.RoleCreateForm
	if err := c.ShouldBindJSON(&roleCreateForm); err != nil {
		c.JSON(http.StatusOK, result.Error(500, err.Error()))
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	roleService := services.NewRoleService(db)
	roleModel, err := roleService.Create(roleCreateForm)
	if err != nil {
		c.JSON(http.StatusOK, result.Error(500, err.Error()))
	}

	c.JSON(http.StatusOK, result.Success(roleModel))
}

func (ctx *Role) Update(c *gin.Context) {
	result := utils.Result{}
	var roleUpdateForm forms.RoleUpdateForm
	if err := c.ShouldBindJSON(&roleUpdateForm); err != nil {
		c.JSON(http.StatusOK, result.Error(500, err.Error()))
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	roleService := services.NewRoleService(db)
	roleModel, err := roleService.Update(roleUpdateForm)
	if err != nil {
		c.JSON(http.StatusOK, result.Error(500, err.Error()))
	}

	c.JSON(http.StatusOK, result.Success(roleModel))
}

func (ctx *Role) Info(c *gin.Context) {
	result := utils.Result{}
	db := c.MustGet("db").(*gorm.DB)
	roleService := services.NewRoleService(db)
	id := c.PostForm("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, result.Error(500, "请选择菜单"))
	}
	data, err := roleService.GetOneById(newId)
	if err != nil {
		c.JSON(http.StatusOK, result.Error(500, err.Error()))
	}

	c.JSON(http.StatusOK, result.Success(data))
}

func (ctx *Role) Delete(c *gin.Context) {
	result := utils.Result{}
	db := c.MustGet("db").(*gorm.DB)
	roleService := services.NewRoleService(db)
	id := c.PostForm("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, result.Error(500, "请选择菜单"))
	}
	roleModel, err := roleService.GetOneById(newId)
	if err != nil {
		c.JSON(http.StatusOK, result.Error(500, err.Error()))
	}

	err = roleService.Delete(roleModel)
	if err != nil {
		c.JSON(http.StatusOK, result.Error(500, err.Error()))
	}

	c.JSON(http.StatusOK, result.Success(nil))
}
