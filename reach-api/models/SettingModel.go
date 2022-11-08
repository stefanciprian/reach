package models

type SettingModel struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (b *SettingModel) TableName() string {
	return "setting"
}
