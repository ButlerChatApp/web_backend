package services

import (
	"context"
	"log"
	"os"
	"time"

	structs "butler_backend/structs"
	utils "butler_backend/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func generateJWT(userName string) (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Failed to load .env: %v", err)
	}

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	claims := jwt.MapClaims{
		"userName": userName,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 24時間有効
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func SignUp(userName, email, password string) structs.SignUpRes {
	ctx := context.Background()
	client, err := utils.NewFirestoreClient()
	if err != nil {
		return structs.SignUpRes{Message: "Failed to connect to Firestore."}
	}
	defer client.Close()

	// すでに登録されているか確認
	iter := client.Collection("users").Where("email", "==", email).Documents(ctx)
	doc, err := iter.Next()
	if err == nil && doc.Exists() {
		return structs.SignUpRes{Message: "User already exists."}
	}

	// パスワードをハッシュ化
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return structs.SignUpRes{Message: "Failed to hash password: " + err.Error()}
	}

	// Firestore にユーザー情報を保存
	uid := utils.GenerateFirebaseUID(email)
	_, err = client.Collection("users").Doc(uid).Set(ctx, map[string]interface{}{
		"email":           email,
		"hashed_password": hashedPassword,
		"name":            userName,
	})
	if err != nil {
		return structs.SignUpRes{Message: "Failed to save user: " + err.Error()}
	}

	return structs.SignUpRes{
		Uid:      uid,
		UserName: userName,
		Message:  "Successfully signed up.",
	}
}

func SignIn(email, password string) structs.SignInRes {
	ctx := context.Background()
	client, err := utils.NewFirestoreClient()
	if err != nil {
		return structs.SignInRes{Message: "Failed to connect to Firestore."}
	}
	defer client.Close()

	// ユーザー確認
	iter := client.Collection("users").Where("email", "==", email).Documents(ctx)
	doc, err := iter.Next()
	// ユーザーが存在しない場合
	if err != nil || !doc.Exists() {
		return structs.SignInRes{Message: "Invalid credentials."}
	}

	storedPassword := doc.Data()["hashed_password"].(string)
	uid := doc.Ref.ID
	userName := doc.Data()["name"].(string)

	// パスワード認証
	if !utils.CheckPasswordHash(password, storedPassword) {
		return structs.SignInRes{Message: "Invalid credentials."}
	}

	// JWT発行
	token, err := generateJWT(userName)
	if err != nil {
		return structs.SignInRes{
			Message: "Failed to generate token: " + err.Error(),
		}
	}

	return structs.SignInRes{
		Uid:      uid,
		UserName: userName,
		Message:  "Successfully signed in.",
		Token:    token,
	}
}
