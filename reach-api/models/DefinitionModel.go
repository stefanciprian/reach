package models

import "gorm.io/gorm"

type DefinitionModel struct {
	gorm.Model
	Name          string              `json:"name"`
	InputSourceId uint                `json:"input_source_id"`
	InputSource   InputSourceModel    `json:"input_source"`
	OutputSources []OutputSourceModel `gorm:"many2many:definition_outputsource"`
	Transformers  []TransformerModel  `gorm:"many2many:definition_transformer"`
}

func (b *DefinitionModel) TableName() string {
	return "definitions"
}
