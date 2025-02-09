package utils

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

// NewFirestoreClient は Firestore クライアントを作成する
func NewFirestoreClient() (*firestore.Client, error) {
	ctx := context.Background()

	// Firebase アプリの初期化
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Failed to load .env: %v", err)
	}
	service_account_json := os.Getenv("SERVICE_ACCOUNT_JSON")
	sa := option.WithCredentialsFile(service_account_json)
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
