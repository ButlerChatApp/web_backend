package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/vertexai/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

// Vertex AI のクライアントを初期化する
func CreateVertexAIClient() (*genai.Client, error) {
	ctx := context.Background()
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Failed to load .env: %v", err)
	}

	service_account_json := os.Getenv("SERVICE_ACCOUNT_JSON")
	projectId := os.Getenv("PROJECT_ID")

	client, err := genai.NewClient(ctx, projectId, "us-central1", option.WithCredentialsFile(service_account_json))
	if err != nil {
		return nil, fmt.Errorf("failed to create Vertex AI client: %v", err)
	}

	return client, nil
}
