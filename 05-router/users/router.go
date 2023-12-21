package users

import "github.com/gin-gonic/gin"

func Router(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", func(ctx *gin.Context) {
		findUsers(ctx)
	})

	routerGroup.POST("/", func(ctx *gin.Context) {
		createUser(ctx)
	})

	routerGroup.GET("/:id", func(ctx *gin.Context) {
		findUser(ctx)
	})

	routerGroup.PUT("/:id", func(ctx *gin.Context) {
		updateUser(ctx)
	})

	routerGroup.DELETE("/:id", func(ctx *gin.Context) {
		deleteUser(ctx)
	})
}