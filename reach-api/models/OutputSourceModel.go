package models

type OutputSourceModel struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Config string `json:"config"` // JSON
	Type   string `json:"type"`   // Enum API post, MQ, Kafka stream, websocket
}

func (b *OutputSourceModel) TableName() string {
	return "output_source"
}
