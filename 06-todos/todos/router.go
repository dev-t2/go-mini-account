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

	routerGroup.PUT("/:id/content", func(ctx *gin.Context) {
		UpdateContent(ctx)
	})

	routerGroup.PUT("/:id/completion", func(ctx *gin.Context) {
		UpdateCompletion(ctx)
	})
}