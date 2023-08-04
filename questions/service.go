package questions

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
	{Code: "1", Description: "Qual é o meu sabor de sorvete favorito?",
		Options: []Option{{Code: "1", Description: "Unicórnio"}, {Code: "2", Description: "Menta com chocolate"}, {Code: "3", Description: "Pistache"}}},
	{Code: "2", Description: "Qual é o meu filme preferido de todos os tempos?",
		Options: []Option{{Code: "1", Description: "Casa comigo"}, {Code: "2", Description: "Melhor amigo da noiva"}, {Code: "3", Description: "Uma linda mulher"}}},
	{Code: "3", Description: "Qual é o meu destino dos seus sonhos para viajar?",
		Options: []Option{{Code: "1", Description: "Austrália"}, {Code: "2", Description: "Tailândia"}, {Code: "3", Description: "Filipinas"}}},
	{Code: "4", Description: "Qual é o meu prato de comida preferido?",
		Options: []Option{{Code: "1", Description: "Strogonoff"}, {Code: "2", Description: "Macarrão"}, {Code: "3", Description: "Feijoada"}}},
	{Code: "5", Description: "Qual é o meu jogo de tabuleiro ou vídeo game preferido?",
		Options: []Option{{Code: "1", Description: "Need for speed "}, {Code: "2", Description: "Counter strike"}, {Code: "3", Description: "Super Mario"}}},
	{Code: "6", Description: "Qual o estilo de filme que mais odeio?",
		Options: []Option{{Code: "1", Description: "Suspense"}, {Code: "2", Description: "Terror"}, {Code: "3", Description: "Drama"}}},
	{Code: "7", Description: "Quantos irmão tenho?",
		Options: []Option{{Code: "1", Description: "Um"}, {Code: "2", Description: "Dois"}, {Code: "3", Description: "Três"}}},
	{Code: "8", Description: "Qual lugar eu sonho morar?",
		Options: []Option{{Code: "1", Description: "Jericoacoara"}, {Code: "2", Description: "São Miguel dos milagres"}, {Code: "3", Description: "Maragogi"}}},
	{Code: "9", Description: "Em quantas cidades já morei?",
		Options: []Option{{Code: "1", Description: "Duas"}, {Code: "2", Description: "Três"}, {Code: "3", Description: "Quatro"}}},
	{Code: "10", Description: "Qual minha cor favorita?",
		Options: []Option{{Code: "1", Description: "Azul"}, {Code: "2", Description: "Vermelho"}, {Code: "3", Description: "Amarelo"}}},
}

func GetAll() []Question {
	return questions
}
