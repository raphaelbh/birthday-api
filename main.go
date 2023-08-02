package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	service "github.com/raphaelbh/birthday-api/questions"
)

func main() {
	router := gin.Default()
	router.GET("/questions", getQuestions)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, service.GetAll())
}
