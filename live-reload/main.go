package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		value := c.Query("channel")
		fmt.Println(value)

		c.String(200, fmt.Sprintf("Seu canal se chama %s", value))
	})

	router.Run(":8080")
}
