package models

import "gorm.io/gorm"

type CaseModel struct {
	gorm.Model
	Payload      string          `json:"payload"`
	DefinitionId uint            `json:"definition_id"`
	Definition   DefinitionModel `json:"definition"`
	SummaryID    uint            `json:"summary_id"`
	Summary      SummaryModel    `json:"summary"`
}

func (b *CaseModel) TableName() string {
	return "cases"
}
