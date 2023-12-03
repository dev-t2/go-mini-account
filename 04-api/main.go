package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
}

var users = []User{}

func main() {
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.CustomRecovery(func(ctx *gin.Context, err any) {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
	}))

	router.GET("/users", func(ctx *gin.Context) {})

	router.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "Not Found")
	})

	router.Run(":8080")
}