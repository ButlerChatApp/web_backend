package services

import (
	"context"
	"time"

	structs "butler_backend/structs"
	utils "butler_backend/utils"

	"github.com/google/uuid"
)

func ChatCreation(chatType, chatName string, participants []structs.Participant) (structs.ChatCreationRes, error) {
	ctx := context.Background()
	client, err := utils.NewFirestoreClient()
	if err != nil {
		return structs.ChatCreationRes{}, err
	}
	defer client.Close()

	chatId := uuid.New().String()
	chatData := map[string]interface{}{
		"chatId":      chatId,
		"type":        chatType,
		"chatName":    chatName,
		"participants": participants,
		"createdAt":   time.Now(),
		"updatedAt":   time.Now(),
	}

	_, err = client.Collection("chats").Doc(chatId).Set(ctx, chatData)
	if err != nil {
		return structs.ChatCreationRes{}, err
	}

	response := structs.ChatCreationRes{
		ChatId:       chatId,
		Type:         chatType,
		ChatName:     chatName,
		Participants: participants,
	}

	return response, nil
}



func GetAllChats() (structs.GetAllChatsRes, error) {
	ctx := context.Background()
	client, err := utils.NewFirestoreClient()
	if err != nil {
		return structs.GetAllChatsRes{}, err
	}
	defer client.Close()

	chatsCollection := client.Collection("chats")
	docs, err := chatsCollection.Documents(ctx).GetAll()
	if err != nil {
		return structs.GetAllChatsRes{}, err
	}

	var chats []structs.Chat
	for _, doc := range docs {
		var chat structs.Chat
		err := doc.DataTo(&chat)
		if err != nil {
			return structs.GetAllChatsRes{}, err
		}
		chats = append(chats, chat)
	}

	response := structs.GetAllChatsRes{
		Chats: chats,
	}

	return response, nil
}