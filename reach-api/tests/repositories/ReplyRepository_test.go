package tests

import (
	"reach/reach-api/models"
	"reach/reach-api/repositories"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestGetAllReplies(t *testing.T) {
	type args struct {
		reply *[]models.ReplyModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repositories.GetAllReplies(tt.args.reply); (err != nil) != tt.wantErr {
				t.Errorf("GetAllReplies() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateReply(t *testing.T) {
	type args struct {
		reply *models.ReplyModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repositories.CreateReply(tt.args.reply); (err != nil) != tt.wantErr {
				t.Errorf("CreateReply() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetReplyByID(t *testing.T) {
	type args struct {
		reply *models.ReplyModel
		id    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repositories.GetReplyByID(tt.args.reply, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("GetReplyByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateReply(t *testing.T) {
	type args struct {
		reply *models.ReplyModel
		id    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repositories.UpdateReply(tt.args.reply, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UpdateReply() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteReply(t *testing.T) {
	type args struct {
		reply *models.ReplyModel
		id    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repositories.DeleteReply(tt.args.reply, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteReply() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
