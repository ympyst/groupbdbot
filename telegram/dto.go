package telegram

type User struct {
	Id int `json:"id"`
	IsBot bool `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
}

type Chat struct {
	Id int `json:"id"`
	Type string `json:"type"`
	Username string `json:"username"`
}

type Message struct {
	Id   int  `json:"message_id"`
	Timestamp int `json:"date"`
	UserFrom User `json:"from"`
	Chat Chat `json:"chat"`
	Text string `json:"text"`
	Entities []MessageEntity `json:"entities"`
}

type MessageEntity struct {
	Type string `json:"type"`
	Offset int `json:"offset"`
	Length int `json:"length"`
	Url string `json:"url"`
	User User `json:"user"`
}

type Update struct {
	Id   int  `json:"update_id"`
	Message Message `json:"message"`
}

type SendMessageResponse struct {
	Method string `json:"method"`
	ChatId int `json:"chat_id"`
	Text string `json:"text"`
}