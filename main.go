package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	questions "github.com/raphaelbh/birthday-api/questions"
	quiz "github.com/raphaelbh/birthday-api/quiz"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/questions", getQuestions)
	router.POST("/quiz", postQuiz)
	router.GET("/rank", getRank)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	router.Run(":" + port)
}

func getQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, questions.GetAll())
}

func getRank(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, quiz.GetRank())
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

	quiz.Create(user, answers)
	c.IndentedJSON(http.StatusCreated, answers)
}
