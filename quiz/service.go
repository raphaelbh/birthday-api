package quiz

import (
	"encoding/json"
	"fmt"
	"sort"

	database "github.com/raphaelbh/birthday-api/database"
	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	User  string
	Score int
	Data  string
}

type Rank struct {
	Position int    `json:"position"`
	User     string `json:"user"`
	Score    int    `json:"score"`
}

var questionsAnswers = map[string]string{
	"1":  "1",
	"2":  "3",
	"3":  "2",
	"4":  "2",
	"5":  "1",
	"6":  "2",
	"7":  "2",
	"8":  "1",
	"9":  "1",
	"10": "3",
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

func getAll() []Quiz {
	var listQuiz []Quiz
	result := db.Find(&listQuiz)
	if result.Error != nil {
		panic("Error getting list of quiz: " + result.Error.Error())
	}
	return listQuiz
}

func GetRank() []Rank {

	var listQuiz = getAll()

	// sort list
	sortFunc := func(i, j int) bool {
		return listQuiz[i].Score > listQuiz[j].Score
	}
	sort.Slice(listQuiz, sortFunc)

	// rank
	var rank []Rank
	var score = -1
	var position = 0
	for _, item := range listQuiz {
		if score != item.Score {
			score = item.Score
			position = len(rank) + 1
		}
		rank = append(rank, Rank{Position: position, User: item.User, Score: item.Score})
	}

	return rank
}
