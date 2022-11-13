package models

import "gorm.io/gorm"

type ReplyModel struct {
	gorm.Model
	Metadata  string       `json:"metadata"`
	CaseId    uint         `json:"case_id"`
	Case      CaseModel    `json:"case"`
	SummaryId uint         `json:"summary_id"`
	Summary   SummaryModel `json:"summary"`
}

func (b *ReplyModel) TableName() string {
	return "replies"
}
