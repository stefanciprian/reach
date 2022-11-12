package tests

import (
	"reach/reach-api/models"
	"testing"
)

func TestInputSourceModel_TableName(t *testing.T) {
	type fields struct {
		Id     uint
		Name   string
		Config string
		Type   string
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
			b := &models.InputSourceModel{
				Id:     tt.fields.Id,
				Name:   tt.fields.Name,
				Config: tt.fields.Config,
				Type:   tt.fields.Type,
			}
			if got := b.TableName(); got != tt.want {
				t.Errorf("InputSourceModel.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
