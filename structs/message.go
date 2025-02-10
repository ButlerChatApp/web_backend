package structs

import "time"

type Message struct {
	MessageId  string    `json:"messageId" binding:"required"`
	ChatId     string    `json:"chatId" binding:"required"`
	SenderId   string    `json:"senderId" binding:"required"`
	SenderName string    `json:"senderName" binding:"required"`
	Content    string    `json:"content" binding:"required"`
	Timestamp  time.Time `json:"timestamp" binding:"required"`
}

type PostMessageReq struct {
	ChatId   string `json:"chatId" binding:"required"`
	SenderId string `json:"senderId" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

type PostMessageRes struct {
	MessageId  string    `json:"messageId" binding:"required"`
	ChatId     string    `json:"chatId" binding:"required"`
	SenderId   string    `json:"senderId" binding:"required"`
	SenderName string    `json:"senderName" binding:"required"`
	Content    string    `json:"content" binding:"required"`
	Timestamp  time.Time `json:"timestamp" binding:"required"`
}

type GetMessagesRes struct {
	Messages []Message `json:"messages" binding:"required"`
}

type EditMessageParam struct {
	MessageId  string `json:"messageId" binding:"required"`
	NewContent string `json:"newContent" binding:"required"`
}

type EditMessageRes struct {
	MessageId  string `json:"chatId" binding:"required"`
	NewContent string `json:"newContent"`
}

type DeleteMessageRes struct {
	MessageId string `json:"messageId" binding:"required"`
}
