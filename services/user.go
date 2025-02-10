package services

import (
	"context"
	"log"

	structs "butler_backend/structs"
	utils "butler_backend/utils"
)

func GetUserName(uid string) (string, error) {
	ctx := context.Background()
	client, err := utils.NewFirestoreClient()
	if err != nil {
		return "", err
	}
	defer client.Close()

	doc, err := client.Collection("users").Doc(uid).Get(ctx)
	if err != nil {
		log.Printf("Failed to get user document: %v", err)
		return "", err
	}

	var user structs.User
	if err := doc.DataTo(&user); err != nil {
		log.Printf("Failed to parse user data: %v", err)
		return "", err
	}

	return user.Name, nil
}
