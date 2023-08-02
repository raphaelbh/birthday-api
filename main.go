package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	service "github.com/raphaelbh/birthday-api/questions"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/questions", getQuestions)
	router.POST("/quiz", postQuiz)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	router.Run(":" + port)
}

func getQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, service.GetAll())
}

func postQuiz(c *gin.Context) {
	var answers map[string]string

	if err := c.BindJSON(&answers); err != nil {
		return
	}

	user := c.GetHeader("x-user")
	if user == "" {
		c.String(http.StatusBadRequest, "Header x-user not informed")
		return
	}

	service.Create(user, answers)
	c.IndentedJSON(http.StatusCreated, answers)
}
