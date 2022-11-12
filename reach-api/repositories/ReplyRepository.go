package repositories

import (
	"reach/reach-api/config"
	"reach/reach-api/models"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllReplies(reply *[]models.ReplyModel) (err error) {
	if err = config.DB.Find(reply).Error; err != nil {
		return err
	}
	return nil
}

func CreateReply(reply *models.ReplyModel) (err error) {
	if err = config.DB.Create(reply).Error; err != nil {
		return err
	}
	return nil
}

func GetReplyByID(reply *models.ReplyModel, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(reply).Error; err != nil {
		return err
	}
	return nil
}

func UpdateReply(reply *models.ReplyModel, id string) (err error) {
	config.DB.Save(reply)
	return nil
}

func DeleteReply(reply *models.ReplyModel, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(reply)
	return nil
}
