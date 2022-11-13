package models

import "gorm.io/gorm"

type InputSourceModel struct {
	gorm.Model
	Name   string `json:"name"`
	Config string `json:"config"` // JSON
	Type   string `json:"type"`   // Enum API post, MQ, Kafka stream, websocket
}

func (b *InputSourceModel) TableName() string {
	return "input_sources"
}
