package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.CustomRecovery(func(ctx *gin.Context, err any) {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
	}))

	router.Static("/", "03-static/public")

	router.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "Not Found")
	})

	router.Run(":8080")
}