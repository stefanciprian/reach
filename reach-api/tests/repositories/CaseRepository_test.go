package repositories

import (
	"reach/reach-api/models"
	"reach/reach-api/repositories"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestGetAllCases(t *testing.T) {
	type args struct {
		caseModel *[]models.CaseModel
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
			if err := repositories.GetAllCases(tt.args.caseModel); (err != nil) != tt.wantErr {
				t.Errorf("GetAllCases() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateCase(t *testing.T) {
	type args struct {
		caseModel *models.CaseModel
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
			if err := repositories.CreateCase(tt.args.caseModel); (err != nil) != tt.wantErr {
				t.Errorf("CreateCase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetCaseByID(t *testing.T) {
	type args struct {
		caseModel *models.CaseModel
		id        string
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
			if err := repositories.GetCaseByID(tt.args.caseModel, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("GetCaseByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateCase(t *testing.T) {
	type args struct {
		caseModel *models.CaseModel
		id        string
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
			if err := repositories.UpdateCase(tt.args.caseModel, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UpdateCase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteCase(t *testing.T) {
	type args struct {
		caseModel *models.CaseModel
		id        string
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
			if err := repositories.DeleteCase(tt.args.caseModel, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
