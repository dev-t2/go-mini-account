package users

import "github.com/gin-gonic/gin"

func Router(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", FindUsers)

	routerGroup.POST("/", CreateUser)

	routerGroup.GET("/:id", FindUser)

	routerGroup.PUT("/:id", UpdateUser)

	routerGroup.DELETE("/:id", DeleteUser)
}