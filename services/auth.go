package services

import (
	"log"
	"os"
	"time"

	structs "butler_backend/structs"
	utils "butler_backend/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var users = make(map[string]string) // email -> hashed password
var userNames = make(map[string]string) // email -> userName

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
	var req structs.SignUpReq

	// すでに登録されているか確認
	if _, exists := users[email]; exists {
		res := structs.SignUpRes{
			Message: "User already exists.",
		}
		return res
	}

	// パスワードをハッシュ化
	hashedPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		res := structs.SignUpRes{
			Message: "Failed to hash password" + err.Error(),
		}
		return res
	}

	// ユーザーを保存
	users[req.Email] = string(hashedPassword)
	userNames[req.Email] = req.UserName

	res := structs.SignUpRes{
		UserName: req.UserName,
		Message: "Successfully signed up.",
	}
	return res
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