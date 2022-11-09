package models

type TransformerModel struct {
	Id uint `json:"id"`
}

func (b *TransformerModel) TableName() string {
	return "transformer"
}
