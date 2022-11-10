package models

type InputSourceModel struct {
	Id uint `json:"id"`
}

func (b *InputSourceModel) TableName() string {
	return "input_source"
}
