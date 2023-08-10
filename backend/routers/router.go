package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/systoms/wegin/backend/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	database := handlers.Database{}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/database/read", database.Read)

	return r
}
