package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	messages "github.com/raphaelbh/birthday-api/messages"
	questions "github.com/raphaelbh/birthday-api/questions"
	quiz "github.com/raphaelbh/birthday-api/quiz"
)

func main() {

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://bailedajack.onrender.com"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Authorization", "Content-Type", "x-user"}

	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(cors.New(config))

	router.GET("/questions", getQuestions)
	router.POST("/quiz", postQuiz)
	router.GET("/rank", getRank)

	router.POST("/messages", postMessage)

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

func postMessage(c *gin.Context) {
	var message messages.Message

	if err := c.BindJSON(&message); err != nil {
		return
	}

	messages.Create(message)
	c.IndentedJSON(http.StatusCreated, message)
}
