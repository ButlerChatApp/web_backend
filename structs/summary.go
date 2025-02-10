package structs

import "time"

type Summary struct {
	SummaryId string    `json:"summaryId" binding:"required"`
	Uid       string    `json:"uid" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	Summary   string    `json:"summary" binding:"required"`
	Timestamp time.Time `json:"timestamp" binding:"required"`
}

type GetSummariesRes struct {
	Summaries []Summary `json:"summaries" binding:"required"`
}

type SummaryRequestReq struct {
	Uid     string `json:"uid" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type SummaryRequestRes struct {
	SummaryId string `json:"summaryId" binding:"required"`
}
