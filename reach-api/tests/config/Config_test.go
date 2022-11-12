package tests

import (
	"reach/reach-api/config"
	"reflect"
	"testing"
)

func TestBuildDBConfig(t *testing.T) {
	tests := []struct {
		name string
		want *config.DBConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := config.BuildDBConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildDBConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDbURL(t *testing.T) {
	type args struct {
		dbConfig *config.DBConfig
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := config.DbURL(tt.args.dbConfig); got != tt.want {
				t.Errorf("DbURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
