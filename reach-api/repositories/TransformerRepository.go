package repositories

import (
	"fmt"

	"reach/reach-api/config"
	"reach/reach-api/models"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllTransformers(transformer *[]models.TransformerModel) (err error) {
	if err = config.DB.Find(transformer).Error; err != nil {
		return err
	}
	return nil
}

func CreateTransformer(transformer *models.TransformerModel) (err error) {
	if err = config.DB.Create(transformer).Error; err != nil {
		return err
	}
	return nil
}

func GetTransformerByID(transformer *models.TransformerModel, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(transformer).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTransformer(transformer *models.TransformerModel, id string) (err error) {
	fmt.Println(transformer)
	config.DB.Save(transformer)
	return nil
}

func DeleteTransformer(transformer *models.TransformerModel, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(transformer)
	return nil
}
