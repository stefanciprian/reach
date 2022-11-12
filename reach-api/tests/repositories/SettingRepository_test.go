package tests

import (
	"reach/reach-api/models"
	"reach/reach-api/repositories"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestGetAllSettings(t *testing.T) {
	type args struct {
		setting *[]models.SettingModel
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
			if err := repositories.GetAllSettings(tt.args.setting); (err != nil) != tt.wantErr {
				t.Errorf("GetAllSettings() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateSetting(t *testing.T) {
	type args struct {
		setting *models.SettingModel
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
			if err := repositories.CreateSetting(tt.args.setting); (err != nil) != tt.wantErr {
				t.Errorf("CreateSetting() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetSettingByID(t *testing.T) {
	type args struct {
		setting *models.SettingModel
		id      string
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
			if err := repositories.GetSettingByID(tt.args.setting, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("GetSettingByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateSetting(t *testing.T) {
	type args struct {
		setting *models.SettingModel
		id      string
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
			if err := repositories.UpdateSetting(tt.args.setting, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UpdateSetting() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteSetting(t *testing.T) {
	type args struct {
		setting *models.SettingModel
		id      string
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
			if err := repositories.DeleteSetting(tt.args.setting, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteSetting() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
