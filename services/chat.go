package services

import (
	"context"
	"log"
	"time"

	structs "butler_backend/structs"
	utils "butler_backend/utils"

	"github.com/google/uuid"
)

func ChatCreation(chatType, chatName, chatCreatorId string, participantsEmails []structs.ParticipantsEmail) (structs.ChatCreationRes, error) {
	ctx := context.Background()
	client, err := utils.NewFirestoreClient()
	if err != nil {
		return structs.ChatCreationRes{}, err
	}
	defer client.Close()

	var participants []structs.Participant
	for i := 0; i < len(participantsEmails); i++ {
		participants = append(participants, structs.Participant{
			Uid: getUid(participantsEmails[i].Email),
		})
	}

	participants = append(participants, structs.Participant{
		Uid: chatCreatorId,
	})

	chatId := uuid.New().String()
	chatData := map[string]interface{}{
		"chatId":       chatId,
		"type":         chatType,
		"chatName":     chatName,
		"participants": participants,
		"createdAt":    time.Now(),
		"updatedAt":    time.Now(),
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

func GetAllChats(uid, selectedType string) (structs.GetAllChatsRes, error) {
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
		data := doc.Data()

		// chatIdとtypeが存在しない場合はスキップ
		chatId, ok := data["chatId"].(string)
		chatType, ok := data["type"].(string)
		if !ok {
			log.Printf("Warning: skipping chat document with invalid or missing chatId: %v", data)
			continue
		}

		// selectedTypeが指定されている場合、一致するもののみを処理
		if selectedType != "" && chatType != selectedType {
			continue
		}

		chat := structs.Chat{
			ChatId: chatId,
			Type:   chatType,
		}

		if participants, ok := data["participants"].([]interface{}); ok {
			for _, p := range participants {
				if participantMap, ok := p.(map[string]interface{}); ok {
					if uidValue, ok := participantMap["Uid"].(string); ok {
						chat.Participants = append(chat.Participants, structs.Participant{
							Uid: uidValue,
						})
					}
				}
			}
		}

		if chatType == "group" {
            if chatName, ok := data["chatName"].(string); ok {
                chat.ChatName = chatName
            }
        } else if chatType == "dm" && contains(chat.Participants, uid) {
			// 相手のuidを取得
			var otherUid string
			for _, participant := range chat.Participants {
				if participant.Uid != uid {
					otherUid = participant.Uid
					break
				}
			}

			// usersコレクションから相手のユーザー名を取得
			userDoc, err := client.Collection("users").Doc(otherUid).Get(ctx)
			if err != nil {
				log.Printf("Warning: failed to get user document for uid %s: %v", otherUid, err)
				continue
			}

			if userName, ok := userDoc.Data()["name"].(string); ok {
				chat.ChatName = userName // ChatNameフィールドに相手のユーザー名を設定
			}
		}

		if contains(chat.Participants, uid) {
			chats = append(chats, chat)
		}
	}

	response := structs.GetAllChatsRes{
		Chats: chats,
	}

	return response, nil
}

func getUid(email string) string {
	ctx := context.Background()
	client, err := utils.NewFirestoreClient()
	if err != nil {
		return ""
	}
	defer client.Close()

	user := client.Collection("users").Where("email", "==", email).Documents(ctx)
	doc, err := user.Next()
	if err != nil {
		return ""
	}
	return doc.Ref.ID
}

// participantsにuidが含まれているかを確認する
func contains(participants []structs.Participant, uid string) bool {
	for _, participant := range participants {
		if participant.Uid == uid {
			return true
		}
	}
	return false
}
