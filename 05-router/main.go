package main

import (
	"net/http"

	"github.com/dev-t2/learn-gin/05-router/users"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()

	engine.Use(gin.Logger())

	engine.Use(gin.CustomRecovery(func(ctx *gin.Context, err any) {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
	}))

	users.Router(engine.Group("/users"))

	engine.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "Not Found")
	})

	engine.Run(":8080")
}