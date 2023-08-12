package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/systoms/wegin/backend/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.New()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	handlerPassport := handlers.Passport{}
	r.POST("/passport/login", handlerPassport.Login)
	r.POST("/passport/logout", handlerPassport.Logout)

	//菜单
	handlerMenu := handlers.Menu{}
	r.POST("/menu/tree", handlerMenu.Tree)
	r.POST("/menu/list", handlerMenu.List)
	r.POST("/menu/create", handlerMenu.Create)
	r.POST("/menu/update", handlerMenu.Update)
	r.POST("/menu/delete", handlerMenu.Delete)
	r.POST("/menu/info", handlerMenu.Info)

	//角色
	handlerRole := handlers.Role{}
	r.POST("/role/list", handlerRole.List)
	r.POST("/role/create", handlerRole.Create)
	r.POST("/role/update", handlerRole.Update)
	r.POST("/role/delete", handlerRole.Delete)
	r.POST("/role/info", handlerRole.Info)

	return r
}
