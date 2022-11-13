package tests

import (
	"reach/reach-api/models"
	"testing"
)

func TestSettingModel_TableName(t *testing.T) {
	type fields struct {
		Id    uint
		Name  string
		Value string
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
			b := &models.SettingModel{
				Name:  tt.fields.Name,
				Value: tt.fields.Value,
			}
			if got := b.TableName(); got != tt.want {
				t.Errorf("SettingModel.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
