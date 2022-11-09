package models

type OutputModel struct {
	Id uint `json:"id"`
}

func (b *OutputModel) TableName() string {
	return "output"
}
