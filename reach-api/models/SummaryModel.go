package models

import "gorm.io/gorm"

type SummaryModel struct {
	gorm.Model
	Replies []ReplyModel `json:"replies"`
}

func (b *SummaryModel) TableName() string {
	return "summaries"
}
