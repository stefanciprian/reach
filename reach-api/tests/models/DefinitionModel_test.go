package tests

import (
	"reach/reach-api/models"
	"testing"
)

func TestDefinitionModel_TableName(t *testing.T) {
	type fields struct {
		Id                uint
		Name              string
		CaseId            uint
		InputSourceModels []models.InputSourceModel
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &models.DefinitionModel{
				Id:                tt.fields.Id,
				Name:              tt.fields.Name,
				CaseId:            tt.fields.CaseId,
				InputSourceModels: tt.fields.InputSourceModels,
			}
			if got := b.TableName(); got != tt.want {
				t.Errorf("DefinitionModel.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
