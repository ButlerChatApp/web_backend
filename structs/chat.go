package structs

type Chat struct {
	ChatId string `json:"chatId" binding:"required"`
	ChatName string `json:"chatName" binding:"required"`
	Type string `json:"type" binding:"required"`
	Participants []Participant `json:"participants" binding:"required"`
}

type Participant struct {
	Uid string `json:"Uid" binding:"required"`
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

type EditMessageRes struct {
	MessageId  string `json:"chatId" binding:"required"`
	NewContent string `json:"newContent"`
}
