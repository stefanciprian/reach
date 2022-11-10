package models

type CaseModel struct {
	Id           uint   `json:"id"`
	Payload      string `json:"payload"`
	DefinitionId uint   `json:"definitionId"`
}

func (b *CaseModel) TableName() string {
	return "case"
}
