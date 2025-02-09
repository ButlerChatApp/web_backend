package structs

type User struct {
	Email string `json:"email" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type GetUserNameReq struct {
	Uid string `json:"uid" binding:"required"`
}