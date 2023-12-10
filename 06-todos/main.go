package main

import (
	"net/http"

	"github.com/dev-t2/learn-gin/06-todos/todos"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()

	engine.Static("/public", "/public")

	engine.Use(gin.Logger())

	engine.Use(gin.CustomRecovery(func(ctx *gin.Context, err any) {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
	}))

	todos.Router(engine.Group("/todos"))

	engine.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "Not Found")
	})

	engine.Run(":8080")
}