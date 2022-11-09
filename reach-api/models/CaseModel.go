package models

type CaseModel struct {
	Id uint `json:"id"`
}

func (b *CaseModel) TableName() string {
	return "case"
}
