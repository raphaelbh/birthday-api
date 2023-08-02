package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	service "github.com/raphaelbh/birthday-api/questions"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/questions", getQuestions)
	router.POST("/quiz", postQuiz)

	router.Run(":80")
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
