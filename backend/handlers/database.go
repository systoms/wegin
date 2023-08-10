package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Database struct {
}

func (ctx *Database) Read(c *gin.Context) {

	c.JSON(http.StatusOK, struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Code:    200,
		Message: "SUCCESS",
	})
}
