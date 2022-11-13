package models

import "gorm.io/gorm"

type TransformerModel struct {
	gorm.Model
}

func (b *TransformerModel) TableName() string {
	return "transformers"
}
