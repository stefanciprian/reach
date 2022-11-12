package tests

import (
	"reach/reach-api/config"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
)

func TestConnectAws(t *testing.T) {
	tests := []struct {
		name string
		want *session.Session
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := config.ConnectAws(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConnectAws() = %v, want %v", got, tt.want)
			}
		})
	}
}
