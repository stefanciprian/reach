package tests

import (
	"reach/reach-api/models"
	"testing"
)

func TestCaseModel_TableName(t *testing.T) {
	type fields struct {
		Id         uint
		Payload    string
		Definition models.DefinitionModel
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
			b := &models.CaseModel{
				Payload:    tt.fields.Payload,
				Definition: tt.fields.Definition,
			}
			if got := b.TableName(); got != tt.want {
				t.Errorf("CaseModel.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
