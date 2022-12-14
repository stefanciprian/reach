package tests

import (
	"reach/reach-api/models"
	"testing"
)

func TestCaseModel_TableName(t *testing.T) {
	type fields struct {
		Id           uint
		Payload      string
		DefinitionId uint
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
				Id:           tt.fields.Id,
				Payload:      tt.fields.Payload,
				DefinitionId: tt.fields.DefinitionId,
			}
			if got := b.TableName(); got != tt.want {
				t.Errorf("CaseModel.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
