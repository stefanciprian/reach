package repositories

import (
	"fmt"

	"reach/reach-api/config"
	"reach/reach-api/models"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllOutputSources(outputSource *[]models.OutputSourceModel) (err error) {
	if err = config.DB.Find(outputSource).Error; err != nil {
		return err
	}
	return nil
}

func CreateOutputSource(outputSource *models.OutputSourceModel) (err error) {
	if err = config.DB.Create(outputSource).Error; err != nil {
		return err
	}
	return nil
}

func GetOutputSourceByID(outputSource *models.OutputSourceModel, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(outputSource).Error; err != nil {
		return err
	}
	return nil
}

func UpdateOutputSource(outputSource *models.OutputSourceModel, id string) (err error) {
	fmt.Println(outputSource)
	config.DB.Save(outputSource)
	return nil
}

func DeleteOutputSource(outputSource *models.OutputSourceModel, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(outputSource)
	return nil
}
