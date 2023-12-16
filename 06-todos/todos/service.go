package todos

import (
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

	todo := Todo{Id: id, Content: strings.Trim(body.Content, " ")}

	todos = append(todos, todo)

	ctx.JSON(http.StatusCreated, todo)
}

func DeleteTodos(ctx *gin.Context) {

}

func UpdateTodo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad Request")

		return
	}

	var body struct { 
		Content string; 
		IsComplete bool;
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.String(http.StatusBadRequest, "Bad Request")

		return
	}

	for index, todo := range todos {
		if todo.Id == id {
			todos[index] = Todo{Id: todo.Id, Content: strings.Trim(body.Content, " "), IsComplete: body.IsComplete}

			ctx.Status(http.StatusNoContent)

			return
		}
	}

	ctx.String(http.StatusNotFound, "Not Found")
}