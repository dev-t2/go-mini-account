package todos

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

var todos = []Todo{}

func findTodos(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"todos": todos})
}

func createTodo(ctx *gin.Context) {
	var body struct { Content string }

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.String(http.StatusBadRequest, "Bad Request")

		return
	}

	var id = 1
	
	if len(todos) > 0 {
		id = todos[len(todos) - 1].Id + 1
	}

	var maxOrder = -1

	for _, todo := range todos {
		if maxOrder < todo.Order {
			maxOrder = todo.Order
		}
	}

	todo := Todo{Id: id, Content: strings.Trim(body.Content, " "), Order: maxOrder + 1}

	todos = append(todos, todo)

	ctx.JSON(http.StatusCreated, todo)
}

func deleteTodos(ctx *gin.Context) {
	ids := strings.Split(ctx.Query("ids"), ",")

	if len(ids) == 0 || strings.Trim(ids[0], " ") == "" {
		todos = []Todo{}

		ctx.Status(http.StatusNoContent)

		return
	}

	deleteIds := []int{}

	for _, todo := range todos {
		if slices.Contains(ids, strconv.Itoa(todo.Id)) {
			deleteIds = append(deleteIds, todo.Id)
		}
	}

	if len(ids) != len(deleteIds) {
		ctx.String(http.StatusBadRequest, "Bad Request")

		return
	}

	deletedTodos := []Todo{}

	for _, todo := range todos {
		if !slices.Contains(deleteIds, todo.Id) {
			deletedTodos = append(deletedTodos, todo)
		}
	}

	todos = deletedTodos

	ctx.Status(http.StatusNoContent)
}

func updateCompletion(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad Request")

		return
	}

	var body struct { IsComplete bool }

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.String(http.StatusBadRequest, "Bad Request")

		return
	}

	for index, todo := range todos {
		if todo.Id == id {
			todos[index].IsComplete = body.IsComplete

			ctx.Status(http.StatusNoContent)

			return
		}
	}

	ctx.String(http.StatusNotFound, "Not Found")
}

func updateContent(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad Request")

		return
	}

	var body struct { Content string }

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.String(http.StatusBadRequest, "Bad Request")

		return
	}

	for index, todo := range todos {
		if todo.Id == id {
			todos[index].Content = strings.Trim(body.Content, " ")

			ctx.Status(http.StatusNoContent)

			return
		}
	}

	ctx.String(http.StatusNotFound, "Not Found")
}