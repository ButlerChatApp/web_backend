package services

import (
	"log"
	"os"
	"time"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	structs "butler_backend/structs"
	utils "butler_backend/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var users = make(map[string]string) // email -> hashed password
var userNames = make(map[string]string) // email -> userName

func generateFirebaseUID(email string) string {
	hash := sha256.New()
	hash.Write([]byte(strings.ToLower(email)))
	return hex.EncodeToString(hash.Sum(nil))
}

func generateJWT(userName string) (string, error) {
	if err:= godotenv.Load(".env"); err != nil {
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
		return structs.SignUpRes{Message: "Failed to connect to Firestore"}
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
	uid := generateFirebaseUID(email)
	_, err = client.Collection("users").Doc(uid).Set(ctx, map[string]interface{}{
		"email":          email,
		"hashed_password": hashedPassword,
		"name":           userName,
	})
	if err != nil {
		return structs.SignUpRes{Message: "Failed to save user: " + err.Error()}
	}

	return structs.SignUpRes{
		UserName: userName,
		Message:  "Successfully signed up.",
	}
}

func SignIn(email, password string) structs.SignInRes {
	var req structs.SignInReq

	// ユーザー確認
	storedPassword, exists := users[req.Email]
	if !exists {
		res := structs.SignInRes{
			Message: "Invalid credentials.",
		}
		return res
	}

	// パスワード認証
	if !utils.CheckPasswordHash(req.Password, storedPassword) {
		res := structs.SignInRes{
			Message: "Invalid credentials.",
		}
		return res
	}

	// JWTトークン発行
	token, err := generateJWT(userNames[req.Email])
	if err != nil {
		res := structs.SignInRes{
			Message: "Failed to generate token:" + err.Error(),
		}
		return res
	}

	res := structs.SignInRes{
		UserName: userNames[req.Email],
		Message: "Successfully signed in.",
		Token: token,
	}
	return res
}