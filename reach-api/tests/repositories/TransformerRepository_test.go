package tests

import (
	"reach/reach-api/models"
	"reach/reach-api/repositories"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestGetAllTransformers(t *testing.T) {
	type args struct {
		transformer *[]models.TransformerModel
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
			if err := repositories.GetAllTransformers(tt.args.transformer); (err != nil) != tt.wantErr {
				t.Errorf("GetAllTransformers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateTransformer(t *testing.T) {
	type args struct {
		transformer *models.TransformerModel
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
			if err := repositories.CreateTransformer(tt.args.transformer); (err != nil) != tt.wantErr {
				t.Errorf("CreateTransformer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetTransformerByID(t *testing.T) {
	type args struct {
		transformer *models.TransformerModel
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
			if err := repositories.GetTransformerByID(tt.args.transformer, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("GetTransformerByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateTransformer(t *testing.T) {
	type args struct {
		transformer *models.TransformerModel
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
			if err := repositories.UpdateTransformer(tt.args.transformer, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTransformer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteTransformer(t *testing.T) {
	type args struct {
		transformer *models.TransformerModel
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
			if err := repositories.DeleteTransformer(tt.args.transformer, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTransformer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
