package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
}

var users = []User{}

func main() {
	engine := gin.New()

	engine.Use(gin.Logger())

	engine.Use(gin.CustomRecovery(func(ctx *gin.Context, err any) {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
	}))

	engine.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{ "users" : users })
	})

	engine.GET("/users/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.String(http.StatusBadRequest, "Bad Request")

			return
		}

		for _, user := range users {
			if user.Id == id {
				ctx.JSON(http.StatusOK, user)

				return
			}
		}

		ctx.String(http.StatusNotFound, "Not Found")
	})

	engine.POST("/users", func(ctx *gin.Context) {
		var body struct { Nickname string }

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.String(http.StatusBadRequest, "Bad Request")

			return
		}

		body.Nickname = strings.Trim(body.Nickname, " ")

		if body.Nickname == "" {
			ctx.String(http.StatusBadRequest, "Bad Request")

			return
		}

		for _, user := range users {
			if user.Nickname == body.Nickname {
				ctx.String(http.StatusBadRequest, "Bad Request")

				return
			}
		}
		
		var id = 1
		
		if len(users) > 0 {
			id = users[len(users) - 1].Id + 1
		}

		user := User{Id: id, Nickname: body.Nickname}

		users = append(users, user)

		ctx.JSON(http.StatusCreated, user)
	})

	engine.PUT("/users/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.String(http.StatusBadRequest, "Bad Request")

			return
		}

		var body struct {Nickname string}

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.String(http.StatusBadRequest, "Bad Request")

			return
		}

		for index, user := range users {
			if user.Id == id {
				users[index].Nickname = body.Nickname

				ctx.Status(http.StatusNoContent)

				return
			}
		}

		ctx.String(http.StatusNotFound, "Not Found")
	})

	engine.DELETE("/users/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.String(http.StatusBadRequest, "Bad Request")

			return
		}

		for index, user := range users {
			if user.Id == id {
				users = append(users[:index], users[index+1:]...)

				ctx.Status(http.StatusNoContent)

				return
			}
		}

		ctx.String(http.StatusNotFound, "Not Found")
	})

	engine.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "Not Found")
	})

	engine.Run(":8080")
}