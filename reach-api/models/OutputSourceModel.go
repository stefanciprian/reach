package models

type OutputSourceModel struct {
	Id uint `json:"id"`
}

func (b *OutputSourceModel) TableName() string {
	return "output_source"
}
