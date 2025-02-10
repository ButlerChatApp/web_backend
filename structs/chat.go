package structs

type Chat struct {
	ChatId       string        `json:"chatId" binding:"required"`
	ChatName     string        `json:"chatName" binding:"required"`
	Type         string        `json:"type" binding:"required"`
	Participants []Participant `json:"participants" binding:"required"`
}

type Participant struct {
	Uid string `json:"Uid" binding:"required"`
}

type ParticipantsEmail struct {
	Email string `json:"email" binding:"required"`
}

type ChatCreationReq struct {
	Type               string              `json:"type" binding:"required"`
	ChatName           string              `json:"chatName"`
	ChatCreatorId      string              `json:"chatCreatorId" binding:"required"`
	ParticipantsEmails []ParticipantsEmail `json:"participantsEmails" binding:"required"`
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
