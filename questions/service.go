package service

import (
	"encoding/json"
	"fmt"
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
	{Code: "2", Description: "Quantos estados brasileiros eu conheco?",
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

	// calc score
	score := 0
	for key, value := range answers {
		if questionsAnswers[key] == value {
			score++
		}
	}

	// save on database
	// user, score, answers

	jsonData, err := json.Marshal(answers)
	if err != nil {
		fmt.Println("Erro ao serializar em JSON:", err)
		return
	}

	jsonString := string(jsonData)

	fmt.Println("user = ", user)
	fmt.Println("score = ", score)
	fmt.Println("data = ", jsonString)
}
