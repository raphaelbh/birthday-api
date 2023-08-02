package service

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

func GetAll() []Question {
	return questions
}
