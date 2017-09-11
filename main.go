package main

import (
	"fmt"
	"time"
)

var tSub = func(now, start time.Time) time.Duration {
	return now.Sub(start)
}

/**
* `Intersection` receives the duration in minutes to show the status.
 */
func Intersection(duration int) chan string {
	NSLights := NewTrafficLight(Green)
	EWLights := NewTrafficLight(Red)
	ticker := time.NewTicker(time.Second * 1)
	var start = time.Now()
	status := make(chan string)
	go func() {
		for t := range ticker.C {
			elapsed := tSub(t, start)
			minutes := int(elapsed.Minutes()) % 60
			seconds := int(elapsed.Seconds()) % 60
			if minutes >= duration {
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
	status := Intersection(30)
	for s := range status {
		fmt.Print("\r", s)
	}
}
