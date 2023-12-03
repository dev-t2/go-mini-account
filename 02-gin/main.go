package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const addr = ":8080"

func main() {
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.CustomRecovery(func(c *gin.Context, err any) {
		c.String(http.StatusInternalServerError, "Internal Server Error")
	}))

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Gin")
	})

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})

	err := router.Run(addr)

	if err != nil {
		log.Fatal(err)
	}
}