package models

type Settings struct {
	ID             string `json:"id" bson:"_id,omitempty"`
	TelegramToken  string `json:"telegramToken" bson:"telegramtoken"`
	TelegramChatId string `json:"telegramChatId" bson:"telegramchatid"`
}
