package tests

import (
	"reach/reach-api/models"
	"testing"
)

func TestTransformerModel_TableName(t *testing.T) {
	type fields struct {
		Id uint
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
			b := &models.TransformerModel{}
			if got := b.TableName(); got != tt.want {
				t.Errorf("TransformerModel.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
