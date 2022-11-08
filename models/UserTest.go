package models

import (
	"fmt"

	"github.com/stefanciprian/reach/config"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllUserTests Fetch all userTest data
func GetAllUserTests(user *[]UserTest) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

//CreateUserTest ... Insert New data
func CreateUserTest(userTest *UserTest) (err error) {
	if err = config.DB.Create(userTest).Error; err != nil {
		return err
	}
	return nil
}

//GetUserTestByID ... Fetch only one userTest by Id
func GetUserTestByID(userTest *UserTest, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(userTest).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUserTest ... Update userTest
func UpdateUserTest(userTest *UserTest, id string) (err error) {
	fmt.Println(userTest)
	config.DB.Save(userTest)
	return nil
}

//DeleteUserTest ... Delete userTest
func DeleteUserTest(userTest *UserTest, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(userTest)
	return nil
}
