package repositories

import (
	"reach/reach-api/config"
	"reach/reach-api/models"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllUsers(user *[]models.UserModel) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(userTest *models.UserModel) (err error) {
	if err = config.DB.Create(userTest).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByID(user *models.UserModel, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *models.UserModel, id string) (err error) {
	config.DB.Save(user)
	return nil
}

func DeleteUser(user *models.UserModel, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(user)
	return nil
}
