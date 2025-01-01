package models

type Chapter struct {
	Number uint   `json:"number"`
	Name   string `json:"name"`
	Pages  []Page `json:"pages"`
}
