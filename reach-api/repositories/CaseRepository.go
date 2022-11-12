package repositories

import (
	"fmt"

	"reach/reach-api/config"
	"reach/reach-api/models"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllCases(caseModel *[]models.CaseModel) (err error) {
	if err = config.DB.Find(caseModel).Error; err != nil {
		return err
	}
	return nil
}

func CreateCase(caseModel *models.CaseModel) (err error) {
	if err = config.DB.Create(caseModel).Error; err != nil {
		return err
	}
	return nil
}

func GetCaseByID(caseModel *models.CaseModel, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(caseModel).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCase(caseModel *models.CaseModel, id string) (err error) {
	fmt.Println(caseModel)
	config.DB.Save(caseModel)
	return nil
}

func DeleteCase(caseModel *models.CaseModel, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(caseModel)
	return nil
}
