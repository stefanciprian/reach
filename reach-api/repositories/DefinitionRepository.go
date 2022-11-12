package repositories

import (
	"reach/reach-api/config"
	"reach/reach-api/models"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllDefinitions(definition *[]models.DefinitionModel) (err error) {
	if err = config.DB.Find(definition).Error; err != nil {
		return err
	}
	return nil
}

func CreateDefinition(definition *models.DefinitionModel) (err error) {
	if err = config.DB.Create(definition).Error; err != nil {
		return err
	}
	return nil
}

func GetDefinitionByID(definition *models.DefinitionModel, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(definition).Error; err != nil {
		return err
	}
	return nil
}

func UpdateDefinition(definition *models.DefinitionModel, id string) (err error) {
	config.DB.Save(definition)
	return nil
}

func DeleteDefinition(definition *models.DefinitionModel, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(definition)
	return nil
}
