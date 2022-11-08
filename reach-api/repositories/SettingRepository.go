package repositories

import (
	"fmt"

	"github.com/stefanciprian/reach/reach-api/config"
	"github.com/stefanciprian/reach/reach-api/models"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllSettings(setting *[]models.SettingModel) (err error) {
	if err = config.DB.Find(setting).Error; err != nil {
		return err
	}
	return nil
}

func CreateSetting(setting *models.SettingModel) (err error) {
	if err = config.DB.Create(setting).Error; err != nil {
		return err
	}
	return nil
}

func GetSettingByID(setting *models.SettingModel, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(setting).Error; err != nil {
		return err
	}
	return nil
}

func UpdateSetting(setting *models.SettingModel, id string) (err error) {
	fmt.Println(setting)
	config.DB.Save(setting)
	return nil
}

func DeleteSetting(setting *models.SettingModel, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(setting)
	return nil
}
