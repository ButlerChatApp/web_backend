package structs

type SignUpReq struct {
	UserName string `json:"userName" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpRes struct {
	UserName string `json:"userName"`
	Message  string `json:"message" binding:"required"`
}

type SignInReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInRes struct {
	UserName string `json:"userName"`
	Message  string `json:"message" binding:"required"`
	Token    string `json:"token"`
}
