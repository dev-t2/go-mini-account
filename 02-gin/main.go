package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()

	engine.Use(gin.Logger())

	engine.Use(gin.CustomRecovery(func(ctx *gin.Context, err any) {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
	}))

	engine.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello Gin")
	})

	engine.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "Not Found")
	})

	engine.Run(":8080")
}