package users

import "github.com/gin-gonic/gin"

func Router(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", func(ctx *gin.Context) {
		FindUsers(ctx)
	})

	routerGroup.POST("/", func(ctx *gin.Context) {
		CreateUser(ctx)
	})

	routerGroup.GET("/:id", func(ctx *gin.Context) {
		FindUser(ctx)
	})

	routerGroup.PUT("/:id", func(ctx *gin.Context) {
		UpdateUser(ctx)
	})

	routerGroup.DELETE("/:id", func(ctx *gin.Context) {
		DeleteUser(ctx)
	})
}