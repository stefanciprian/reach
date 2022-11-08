package models

type Setting struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (b *Setting) TableName() string {
	return "setting"
}
