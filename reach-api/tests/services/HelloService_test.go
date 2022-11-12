package tests

import (
	"reach/reach-api/services"
	"testing"
)

func TestHelloService(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := services.HelloService(); got != tt.want {
				t.Errorf("HelloService() = %v, want %v", got, tt.want)
			}
		})
	}
}
