package main

import (
	"fmt"
	"sync"

	"github.com/double_booked/models"
	"github.com/double_booked/utils"
)

func main() {
	var wg sync.WaitGroup
	var events []models.Event
	events = []models.Event{
		models.Event{"Event 1", 8, 9, 1, },
		models.Event{"Event 7", 8, 9, 1, },
		models.Event{"Event 2", 8, 9, 2, },
		models.Event{"Event 3", 8, 9, 3, },
		models.Event{"Event 4", 8, 9, 4, },
		models.Event{"Event 5", 8, 9, 5, },
		models.Event{"Event 6", 8, 9, 5, },
	}
	fmt.Println("Overlaping events:")
	wg.Add(1)
	go utils.GetOverlappingEvents(events, &wg)
	wg.Wait()
}
