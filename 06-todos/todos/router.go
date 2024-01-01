package todos

import (
	"github.com/gin-gonic/gin"
)

func Router(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", func(ctx *gin.Context) {
		findTodos(ctx)
	})

	routerGroup.POST("/", func(ctx *gin.Context) {
		createTodo(ctx)
	})

	routerGroup.DELETE("/", func(ctx *gin.Context) {
		deleteTodos(ctx)
	})

	routerGroup.PATCH("/:id/completion", func(ctx *gin.Context) {
		updateCompletion(ctx)
	})

	routerGroup.PATCH("/:id/content", func(ctx *gin.Context) {
		updateContent(ctx)
	})
}