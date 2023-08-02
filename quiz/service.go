package quiz

import (
	"encoding/json"
	"fmt"

	database "github.com/raphaelbh/birthday-api/database"
	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	User  string
	Score int
	Data  string
}

var questionsAnswers = map[string]string{
	"1": "2",
	"2": "2",
}

var db = database.Connect()

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

	err = db.AutoMigrate(&Quiz{})
	if err != nil {
		panic("Error in table creation: " + err.Error())
	}

	newQuiz := Quiz{User: user, Score: score, Data: data}

	result := db.Create(&newQuiz)
	if result.Error != nil {
		panic("Erro saving quiz: " + result.Error.Error())
	}

	fmt.Println("Quiz saved")
}

func GetAll() []Quiz {
	var listQuiz []Quiz
	result := db.Find(&listQuiz)
	if result.Error != nil {
		panic("Error getting list of quiz: " + result.Error.Error())
	}
	return listQuiz
}
