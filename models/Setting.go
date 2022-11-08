package models

import (
	"fmt"

	"github.com/stefanciprian/reach/config"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllSettings Fetch all setting data
func GetAllSettings(setting *[]Setting) (err error) {
	if err = config.DB.Find(setting).Error; err != nil {
		return err
	}
	return nil
}

//CreateSetting ... Insert New Setting
func CreateSetting(setting *Setting) (err error) {
	if err = config.DB.Create(setting).Error; err != nil {
		return err
	}
	return nil
}

//GetSettingByID ... Fetch only one Setting by Id
func GetSettingByID(setting *Setting, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(setting).Error; err != nil {
		return err
	}
	return nil
}

//UpdateSetting ... Update setting
func UpdateSetting(setting *Setting, id string) (err error) {
	fmt.Println(setting)
	config.DB.Save(setting)
	return nil
}

//DeleteSetting ... Delete setting
func DeleteSetting(setting *Setting, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(setting)
	return nil
}
