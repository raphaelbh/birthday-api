package service

import (
	"encoding/json"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Option struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Question struct {
	Code        string   `json:"code"`
	Description string   `json:"description"`
	Options     []Option `json:"options"`
}

var questions = []Question{
	{Code: "1", Description: "Quantos irmãos você tem?",
		Options: []Option{{Code: "1", Description: "1"}, {Code: "2", Description: "2"}, {Code: "3", Description: "3"}}},
	{Code: "2", Description: "Quantos estados brasileiros eu conheço?",
		Options: []Option{{Code: "1", Description: "3"}, {Code: "2", Description: "6"}, {Code: "3", Description: "9"}}},
}

var questionsAnswers = map[string]string{
	"1": "2",
	"2": "2",
}

func GetAll() []Question {
	return questions
}

func Create(user string, answers map[string]string) {

	score := 0
	for key, value := range answers {
		if questionsAnswers[key] == value {
			score++
		}
	}

	jsonData, err := json.Marshal(answers)
	if err != nil {
		fmt.Println("Error in json serialization:", err)
		return
	}

	data := string(jsonData)

	type Quiz struct {
		gorm.Model
		User  string
		Score int
		Data  string
	}

	dsn, _ := os.LookupEnv("INTERNAL_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error in database connection: " + err.Error())
	}

	// Criar a tabela se ela ainda não existir (opcional)
	err = db.AutoMigrate(&Quiz{})
	if err != nil {
		panic("Error in table creation: " + err.Error())
	}

	// Criar um novo usuário
	newQuiz := Quiz{User: user, Score: score, Data: data}

	// Salvar o usuário no banco de dados
	result := db.Create(&newQuiz)
	if result.Error != nil {
		panic("Erro saving quiz: " + result.Error.Error())
	}

	fmt.Println("Quiz saved")
}
