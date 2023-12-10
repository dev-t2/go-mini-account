package todos

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var todos = []Todo{}

func FindTodos(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"todos": todos})
}

func CreateTodo(ctx *gin.Context) {
	var body struct { Content string }

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.String(http.StatusBadRequest, "Bad Request")

		return
	}

	var id = 1
	
	if len(todos) > 0 {
		id = todos[len(todos) - 1].Id + 1
	}

	todo := Todo{Id: id, Content: strings.Trim(body.Content, " "), IsComplete: false}

	todos = append(todos, todo)

	ctx.JSON(http.StatusCreated, todo)
}

func DeleteTodos(ctx *gin.Context) {

}

func UpdateTodo(ctx *gin.Context) {
	_, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad Request")

		return
	}

	var body struct { Content string; IsComplete bool }

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.String(http.StatusBadRequest, "Bad Request")

		return
	}

	fmt.Println(body)
}