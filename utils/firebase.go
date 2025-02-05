package utils

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// NewFirestoreClient は Firestore クライアントを作成する
func NewFirestoreClient() (*firestore.Client, error) {
	ctx := context.Background()

	// Firebase アプリの初期化
	sa := option.WithCredentialsFile("firebase_service_account.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Printf("Failed to initialize Firebase: %v", err)
		return nil, err
	}

	// Firestore クライアントの作成
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Printf("Failed to create Firestore client: %v", err)
		return nil, err
	}

	log.Println("Firestore client created successfully.")
	return client, nil
}

func GenerateFirebaseUID(email string) string {
	hash := sha256.New()
	hash.Write([]byte(strings.ToLower(email)))
	return hex.EncodeToString(hash.Sum(nil))
}
