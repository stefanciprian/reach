package tests

import (
	"reach/reach-api/schedule"
	"testing"
)

func TestCollectDataSchedule_Run(t *testing.T) {
	tests := []struct {
		name string
		e    schedule.CollectDataSchedule
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := schedule.CollectDataSchedule{}
			e.Run()
		})
	}
}
