package tests

import (
	"reach/reach-api/models"
	"reach/reach-api/repositories"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestGetAllOutputSources(t *testing.T) {
	type args struct {
		outputSource *[]models.OutputSourceModel
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
			if err := repositories.GetAllOutputSources(tt.args.outputSource); (err != nil) != tt.wantErr {
				t.Errorf("GetAllOutputSources() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateOutputSource(t *testing.T) {
	type args struct {
		outputSource *models.OutputSourceModel
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
			if err := repositories.CreateOutputSource(tt.args.outputSource); (err != nil) != tt.wantErr {
				t.Errorf("CreateOutputSource() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetOutputSourceByID(t *testing.T) {
	type args struct {
		outputSource *models.OutputSourceModel
		id           string
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
			if err := repositories.GetOutputSourceByID(tt.args.outputSource, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("GetOutputSourceByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateOutputSource(t *testing.T) {
	type args struct {
		outputSource *models.OutputSourceModel
		id           string
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
			if err := repositories.UpdateOutputSource(tt.args.outputSource, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UpdateOutputSource() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteOutputSource(t *testing.T) {
	type args struct {
		outputSource *models.OutputSourceModel
		id           string
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
			if err := repositories.DeleteOutputSource(tt.args.outputSource, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteOutputSource() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
