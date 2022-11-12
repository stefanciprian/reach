package tests

import (
	"reach/reach-api/models"
	"reach/reach-api/repositories"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestGetAllUsers(t *testing.T) {
	type args struct {
		user *[]models.UserModel
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
			if err := repositories.GetAllUsers(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("GetAllUsers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	type args struct {
		userTest *models.UserModel
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
			if err := repositories.CreateUser(tt.args.userTest); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUserByID(t *testing.T) {
	type args struct {
		user *models.UserModel
		id   string
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
			if err := repositories.GetUserByID(tt.args.user, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	type args struct {
		user *models.UserModel
		id   string
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
			if err := repositories.UpdateUser(tt.args.user, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	type args struct {
		user *models.UserModel
		id   string
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
			if err := repositories.DeleteUser(tt.args.user, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
