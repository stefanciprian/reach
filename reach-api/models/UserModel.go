package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (b *UserModel) TableName() string {
	return "users"
}
