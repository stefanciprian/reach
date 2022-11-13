package models

import "gorm.io/gorm"

type OutputSourceModel struct {
	gorm.Model
	Name   string `json:"name"`
	Config string `json:"config"` // JSON
	Type   string `json:"type"`   // Enum API post, MQ, Kafka stream, websocket
}

func (b *OutputSourceModel) TableName() string {
	return "output_sources"
}
