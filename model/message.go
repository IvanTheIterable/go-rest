package model

type Message struct {
	Id      string `json:"id" `
	Date    int64  `json:"date" `
	Payload string `json:"payload" `
}
