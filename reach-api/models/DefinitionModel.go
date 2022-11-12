package models

type DefinitionModel struct {
	Id                uint               `json:"id"`
	Name              string             `json:"name"`
	CaseId            uint               `json:"case_id"`
	InputSourceModels []InputSourceModel `json:"input_source_models"`
}

func (b *DefinitionModel) TableName() string {
	return "definition"
}
