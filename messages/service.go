package messages

import (
	"fmt"

	"github.com/raphaelbh/birthday-api/database"
	"gorm.io/gorm"
)

var db = database.Connect()

type Message struct {
	gorm.Model
	User    string `json:"user"`
	Message string `json:"message" gorm:"type:text"`
}

func Create(message Message) {

	err := db.AutoMigrate(&Message{})
	if err != nil {
		panic("Error in table creation: " + err.Error())
	}

	result := db.Create(&message)
	if result.Error != nil {
		panic("Erro saving message: " + result.Error.Error())
	}

	fmt.Println("Message saved")
}

func GetAll() []Message {
	var messages []Message
	result := db.Find(&messages)
	if result.Error != nil {
		panic("Error getting list of quiz: " + result.Error.Error())
	}
	return messages
}
