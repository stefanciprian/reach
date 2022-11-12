package tests

import (
	"reach/reach-api/models"
	"testing"
)

func TestUserModel_TableName(t *testing.T) {
	type fields struct {
		Id      uint
		Name    string
		Email   string
		Phone   string
		Address string
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
			b := &models.UserModel{
				Id:      tt.fields.Id,
				Name:    tt.fields.Name,
				Email:   tt.fields.Email,
				Phone:   tt.fields.Phone,
				Address: tt.fields.Address,
			}
			if got := b.TableName(); got != tt.want {
				t.Errorf("UserModel.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
