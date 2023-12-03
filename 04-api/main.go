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

	router.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"users": users})
	})

	router.POST("/users", func(ctx *gin.Context) {
		var user User

		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.String(http.StatusBadRequest, "Bad Request")

			return
		}

		var id = 1
		
		if len(users) > 0 {
			id = users[len(users) - 1].ID + 1
		}

		user.ID = id

		users = append(users, user)

		ctx.JSON(http.StatusCreated, user)
	})

	router.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "Not Found")
	})

	router.Run(":8080")
}