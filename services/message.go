package services

import (
	"context"
	"time"

	structs "butler_backend/structs"
	utils "butler_backend/utils"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
)

func GetMessages(chatId string) (structs.GetMessagesRes, error) {
	ctx := context.Background()
	client, err := utils.NewFirestoreClient()
	if err != nil {
		return structs.GetMessagesRes{}, err
	}
	defer client.Close()

	messagesCollection := client.Collection("messages").Where("chatId", "==", chatId).OrderBy("timestamp", firestore.Asc)
	docs, err := messagesCollection.Documents(ctx).GetAll()
	if err != nil {
		return structs.GetMessagesRes{}, err
	}

	var messages []structs.Message
	for _, doc := range docs {
		var message structs.Message
		err := doc.DataTo(&message)
		if err != nil {
			return structs.GetMessagesRes{}, err
		}
		messages = append(messages, message)
	}

	response := structs.GetMessagesRes{
		Messages: messages,
	}

	return response, nil
}

func PostMessage(chatId, senderId, content string) (structs.PostMessageRes, error) {
	ctx := context.Background()
	client, err := utils.NewFirestoreClient()
	if err != nil {
		return structs.PostMessageRes{}, err
	}
	defer client.Close()

	messageId := uuid.New().String()
	jst, _ := time.LoadLocation("Asia/Tokyo")
	timestamp := time.Now().In(jst)
	messageData := map[string]interface{}{
		"messageId": messageId,
		"chatId":    chatId,
		"senderId":  senderId,
		"content":   content,
		"timestamp": timestamp,
	}

	_, err = client.Collection("messages").Doc(messageId).Set(ctx, messageData)
	if err != nil {
		return structs.PostMessageRes{}, err
	}

	response := structs.PostMessageRes{
		MessageId: messageId,
		ChatId:    chatId,
		SenderId:  senderId,
		Content:   content,
		Timestamp: timestamp,
	}

	return response, nil
}

func EditMessage(messageId, newContent string) (structs.EditMessageRes, error) {
	ctx := context.Background()
	client, err := utils.NewFirestoreClient()
	if err != nil {
		return structs.EditMessageRes{}, err
	}
	defer client.Close()

	docRef := client.Collection("messages").Where("messageId", "==", messageId).Limit(1)
	docSnapshot, err := docRef.Documents(ctx).Next()

	if err != nil {
		return structs.EditMessageRes{}, err
	}

	_, err = docSnapshot.Ref.Update(ctx, []firestore.Update{
		{
			Path:  "content",
			Value: newContent,
		},
	})
	if err != nil {
		return structs.EditMessageRes{}, err
	}

	response := structs.EditMessageRes{
		MessageId: messageId,
		NewContent: newContent,
	}
	
	return response, nil
}

func DeleteMessage(messageId string) (structs.DeleteMessageRes, error) {
	ctx := context.Background()
	client, err := utils.NewFirestoreClient()
	if err != nil {
		return structs.DeleteMessageRes{}, err
	}
	defer client.Close()

	docRef := client.Collection("messages").Where("messageId", "==", messageId).Limit(1)
	docs, err := docRef.Documents(ctx).GetAll()
	if err != nil || len(docs) == 0 {
		return structs.DeleteMessageRes{}, err
	}

	_, err = docs[0].Ref.Delete(ctx)
	if err != nil {
		return structs.DeleteMessageRes{}, err
	}

	response := structs.DeleteMessageRes{
		MessageId: messageId,
	}

	return response, nil
}