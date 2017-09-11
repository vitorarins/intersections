package main

import (
	"fmt"
	"time"
)

var colorDurations = []time.Duration{
	(4 * time.Minute) + (30 * time.Second), // Green
	30 * time.Second,                       // Yellow
	5 * time.Minute,                        // Red
}

/**
* `Intersection` receives the duration in minutes to show the status.
 */
func Intersection(duration time.Duration, colorDurations []time.Duration) chan string {
	NSLights := NewTrafficLight(Green, colorDurations)
	EWLights := NewTrafficLight(Red, colorDurations)
	ticker := time.NewTicker(time.Second * 1)
	var start = time.Now()
	status := make(chan string)
	go func() {
		for t := range ticker.C {
			elapsed := t.Sub(start)
			minutes := int(elapsed.Minutes()) % 60
			seconds := int(elapsed.Seconds()) % 60
			if elapsed >= duration {
				close(status)
				break
			}
			status <- fmt.Sprintf("NS: %s, EW: %s - %02d:%02d",
				NSLights.color,
				EWLights.color,
				minutes,
				seconds)
		}
	}()
	return status
}

func main() {
	status := Intersection(30*time.Minute, colorDurations)
	for s := range status {
		fmt.Print(s, "     \r")
	}
}
