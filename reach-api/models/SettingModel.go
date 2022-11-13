package models

import "gorm.io/gorm"

type SettingModel struct {
	gorm.Model
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (b *SettingModel) TableName() string {
	return "settings"
}
