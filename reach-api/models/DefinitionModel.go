package models

type DefinitionModel struct {
	Id uint `json:"id"`
}

func (b *DefinitionModel) TableName() string {
	return "definition"
}
