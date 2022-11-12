package tests

import (
	"reach/reach-api/models"
	"testing"
)

func TestReplyModel_TableName(t *testing.T) {
	type fields struct {
		Id       uint
		Metadata string
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
			b := &models.ReplyModel{
				Id:       tt.fields.Id,
				Metadata: tt.fields.Metadata,
			}
			if got := b.TableName(); got != tt.want {
				t.Errorf("ReplyModel.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
