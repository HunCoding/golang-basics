package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/test", func(ctx *gin.Context) {
		fmt.Println(ctx.Request.Header)
		ctx.String(200, "ok")
	})

	router.Run(":9091")
}
