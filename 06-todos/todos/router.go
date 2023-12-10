package todos

import "github.com/gin-gonic/gin"

func Router(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", func(ctx *gin.Context) {
		FindTodos(ctx)
	})

	routerGroup.POST("/", func(ctx *gin.Context) {
		CreateTodo(ctx)
	})

	routerGroup.DELETE("/", func(ctx *gin.Context) {
		DeleteTodos(ctx)
	})

	routerGroup.PUT("/:id", func(ctx *gin.Context) {
		UpdateTodo(ctx)
	})
}