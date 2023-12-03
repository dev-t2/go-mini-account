package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const addr = ":8080"

func main() {
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.CustomRecovery(func(ctx *gin.Context, err any) {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
	}))

	router.Static("/", "03-static/public")

	router.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "Not Found")
	})

	fmt.Printf("Server running at http://localhost%s\n", addr)

	err := router.Run(addr)

	if err != nil {
		log.Fatal(err)
	}
}