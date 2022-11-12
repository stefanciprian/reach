package repositories

import (
	"reach/reach-api/models"
	"reach/reach-api/repositories"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestGetAllDefinitions(t *testing.T) {
	type args struct {
		definition *[]models.DefinitionModel
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
			if err := repositories.GetAllDefinitions(tt.args.definition); (err != nil) != tt.wantErr {
				t.Errorf("GetAllDefinitions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateDefinition(t *testing.T) {
	type args struct {
		definition *models.DefinitionModel
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
			if err := repositories.CreateDefinition(tt.args.definition); (err != nil) != tt.wantErr {
				t.Errorf("CreateDefinition() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetDefinitionByID(t *testing.T) {
	type args struct {
		definition *models.DefinitionModel
		id         string
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
			if err := repositories.GetDefinitionByID(tt.args.definition, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("GetDefinitionByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateDefinition(t *testing.T) {
	type args struct {
		definition *models.DefinitionModel
		id         string
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
			if err := repositories.UpdateDefinition(tt.args.definition, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UpdateDefinition() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteDefinition(t *testing.T) {
	type args struct {
		definition *models.DefinitionModel
		id         string
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
			if err := repositories.DeleteDefinition(tt.args.definition, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteDefinition() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
