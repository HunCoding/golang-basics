package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name string `json:"name" validate:"required"`
	Age  int32  `json:"age" validate:"required,min=0,max=130"`
}

func main() {

	r := gin.Default()

	r.POST("/welcome", func(c *gin.Context) {

		var user User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, "Erro ao converter dados")
			return
		}

		validate := validator.New()
		err := validate.Struct(user)

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, "Erro ao validar dados")
			return
		}

		c.JSON(http.StatusOK, user)
		return
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
