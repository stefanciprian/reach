package schedule

import (
	"log"
)

// Job Specific Functions
type CollectDataSchedule struct {
	// filtered
}

func (e CollectDataSchedule) Run() {
	log.Println("I am running schedule - CollectDataSchedule")
}
