package structs

type Chat struct {
	ChatId string `json:"chatId" binding:"required"`
}

type Participant struct {
	Uid string `json:"uid" binding:"required"`
}

type ChatCreationReq struct {
	Type         string        `json:"type" binding:"required"`
	ChatName     string        `json:"chatName"`
	Participants []Participant `json:"participants" binding:"required"`
}

type ChatCreationRes struct {
	ChatId       string        `json:"chatId" binding:"required"`
	Type         string        `json:"type" binding:"required"`
	ChatName     string        `json:"chatName"`
	Participants []Participant `json:"participants" binding:"required"`
}

type GetAllChatsRes struct {
	Chats []Chat `json:"chats" binding:"required"`
}
