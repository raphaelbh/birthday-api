package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	messages "github.com/raphaelbh/birthday-api/messages"
	photos "github.com/raphaelbh/birthday-api/photos"
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
	router.GET("/messages", getMessages)

	router.POST("/photos", postPhoto)
	router.GET("/photos", getPhotos)

	router.GET("/health", health)

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

	user, err := url.QueryUnescape(user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Decodification error"})
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

func getMessages(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, messages.GetAll())
}

func postPhoto(c *gin.Context) {

	fmt.Println("[main.postPhoto] Iniciando processando")

	user := c.GetHeader("x-user")
	if user == "" {
		c.String(http.StatusBadRequest, "Header x-user not informed")
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files := form.File["images[]"]
	fmt.Println("[main.postPhoto] Quantidade de arquivos recebidos: ", len(files))

	if len(files) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Images required"})
		return
	}

	if len(files) > 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Max of 10 images"})
		return
	}

	fmt.Println("[main.postPhoto] Chamando service (photos.Upload())")
	uploadErr := photos.Upload(&user, files)
	if uploadErr != nil {
		fmt.Println("[main.postPhoto] Erro ao executar o servico: ", uploadErr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": uploadErr.Error()})
		return
	}

	c.Writer.WriteHeader(204)
}

func getPhotos(c *gin.Context) {

	photos, err := photos.GetAll()
	if err != nil {
		fmt.Println("[main.postPhoto] Erro ao executar o servico: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, photos)
}

func health(c *gin.Context) {
	c.Writer.WriteHeader(200)
}
