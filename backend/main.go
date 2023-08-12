package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/systoms/wegin/backend/routers"
	"github.com/systoms/wegin/config"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// 注册服务（路由）
func regsiterServer(config *config.Config, db *gorm.DB) *gin.Engine {
	e := routers.SetupRouter()
	e.Use(func(c *gin.Context) {
		c.Set("config", config)
		c.Set("db", db)
		c.Next()
	})
	return e
}

// 开始运行
func Run(config *config.Config, db *gorm.DB) error {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      regsiterServer(config, db),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return server.ListenAndServe()
}
