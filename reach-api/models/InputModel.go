package models

type InputModel struct {
	Id uint `json:"id"`
}

func (b *InputModel) TableName() string {
	return "input"
}
