package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const addr = ":8080"

func main() {
	r := gin.New()

	r.SetTrustedProxies(nil)

	r.Use(gin.Logger())

	r.Use(gin.CustomRecovery(func(c *gin.Context, err any) {
		c.String(http.StatusInternalServerError, "Internal Server Error")
	}))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Gin")
	})

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})

	err := r.Run(addr)

	if err != nil {
		log.Fatal(err)
	}
}