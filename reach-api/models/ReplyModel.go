package models

type ReplyModel struct {
	Id       uint   `json:"id"`
	Metadata string `json:"metadata"`
}

func (b *ReplyModel) TableName() string {
	return "reply"
}
