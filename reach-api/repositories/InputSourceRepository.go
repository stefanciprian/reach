package repositories

import (
	"fmt"

	"reach/reach-api/config"
	"reach/reach-api/models"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllInputSources(inputSource *[]models.InputSourceModel) (err error) {
	if err = config.DB.Find(inputSource).Error; err != nil {
		return err
	}
	return nil
}

func CreateInputSource(inputSource *models.InputSourceModel) (err error) {
	if err = config.DB.Create(inputSource).Error; err != nil {
		return err
	}
	return nil
}

func GetInputSourceByID(inputSource *models.InputSourceModel, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(inputSource).Error; err != nil {
		return err
	}
	return nil
}

func UpdateInputSource(inputSource *models.InputSourceModel, id string) (err error) {
	fmt.Println(inputSource)
	config.DB.Save(inputSource)
	return nil
}

func DeleteInputSource(inputSource *models.InputSourceModel, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(inputSource)
	return nil
}
