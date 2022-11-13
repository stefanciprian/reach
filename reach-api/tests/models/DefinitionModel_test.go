package tests

import (
	"reach/reach-api/models"
	"testing"
)

func TestDefinitionModel_TableName(t *testing.T) {
	type fields struct {
		Id            uint
		Name          string
		InputSource   models.InputSourceModel
		OutputSources []models.OutputSourceModel
		Transformers  []models.TransformerModel
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
				Name:          tt.fields.Name,
				InputSource:   tt.fields.InputSource,
				OutputSources: tt.fields.OutputSources,
				Transformers:  tt.fields.Transformers,
			}
			if got := b.TableName(); got != tt.want {
				t.Errorf("DefinitionModel.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
