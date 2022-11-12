package tests

import (
	"reach/reach-api/models"
	"reach/reach-api/repositories"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestGetAllInputSources(t *testing.T) {
	type args struct {
		inputSource *[]models.InputSourceModel
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
			if err := repositories.GetAllInputSources(tt.args.inputSource); (err != nil) != tt.wantErr {
				t.Errorf("GetAllInputSources() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateInputSource(t *testing.T) {
	type args struct {
		inputSource *models.InputSourceModel
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
			if err := repositories.CreateInputSource(tt.args.inputSource); (err != nil) != tt.wantErr {
				t.Errorf("CreateInputSource() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetInputSourceByID(t *testing.T) {
	type args struct {
		inputSource *models.InputSourceModel
		id          string
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
			if err := repositories.GetInputSourceByID(tt.args.inputSource, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("GetInputSourceByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateInputSource(t *testing.T) {
	type args struct {
		inputSource *models.InputSourceModel
		id          string
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
			if err := repositories.UpdateInputSource(tt.args.inputSource, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UpdateInputSource() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteInputSource(t *testing.T) {
	type args struct {
		inputSource *models.InputSourceModel
		id          string
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
			if err := repositories.DeleteInputSource(tt.args.inputSource, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteInputSource() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
