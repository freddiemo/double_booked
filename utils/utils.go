package utils

import (
	"fmt"
	"sync"

	"github.com/double_booked/models"
)

func GetOverlappingEvents(events []models.Event, wg *sync.WaitGroup) {
	e := events[0]
	for i := 1; i < len(events); i++ {
		if e.Day == events[i].Day && e.Start_time == events[i].Start_time && e.End_time == events[i].End_time {
			fmt.Printf("%s, %s \n", e.Name, events[i].Name)
		}
		e = events[i]
	}
	wg.Done()
}