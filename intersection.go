package main

import (
	"fmt"
	"time"
)

// Intersection is a representation of a virtual road intersection.
// With traffic lights for North, South (NS), East and West (EW)
type Intersection struct {
	nsLights       *TrafficLight
	ewLights       *TrafficLight
	greenLights    *TrafficLight
	redLights      *TrafficLight
	maxDuration    time.Duration
	colorDurations map[string]time.Duration
	outputStatus   chan string
	currentState   string
	start          time.Time
	timer          *time.Timer
}

// NewIntersection allocates a new Intersection that will show it's outputStatus during maxDuration.
// And change colors on the traffic lights based on colorDurations.
func NewIntersection(maxDuration time.Duration, colorDurations map[string]time.Duration) *Intersection {
	nsLights := NewTrafficLight(Green)
	ewLights := NewTrafficLight(Red)
	greenLights := nsLights
	redLights := ewLights
	outputStatus := make(chan string)
	colorDurations["green"] = colorDurations["switch"] - colorDurations["yellow"]
	timer := time.NewTimer(colorDurations["green"])
	intersection := &Intersection{nsLights,
		ewLights,
		greenLights,
		redLights,
		maxDuration,
		colorDurations,
		outputStatus,
		"green",
		time.Now(),
		timer,
	}

	// This is necessary to report the initial status of the lights
	// Calling `go writeStatus()` here can cause a race condition.
	go func() {
		outputStatus <- "(N, S): Green, (E, W): Red - 00:00.00"
	}()
	go intersection.switchLights()

	return intersection
}

func (intersection *Intersection) writeStatus() {
	elapsed := time.Now().Sub(intersection.start)
	minutes := int(elapsed.Minutes()) % 60
	seconds := int(elapsed.Seconds()) % 60
	milliseconds := int((elapsed.Nanoseconds()/1000)/1000) % 1000
	intersection.outputStatus <- fmt.Sprintf("(N, S): %s, (E, W): %s - %02d:%02d.%s",
		intersection.nsLights.color,
		intersection.ewLights.color,
		minutes,
		seconds,
		fmt.Sprintf("%03d", milliseconds)[:2])
}

func (intersection *Intersection) switchLights() {
	<-intersection.timer.C
	switch intersection.currentState {
	case "green":
		intersection.greenLights.switchColor()
		intersection.currentState = "yellow"
		intersection.timer = time.NewTimer(intersection.colorDurations["yellow"])
	case "yellow":
		intersection.nsLights.switchColor()
		intersection.ewLights.switchColor()

		// In order to avoid both green and red lights pointing to the same lights.
		copyOfGreen := intersection.greenLights
		intersection.greenLights = intersection.redLights
		intersection.redLights = copyOfGreen
		intersection.currentState = "green"
		intersection.timer = time.NewTimer(intersection.colorDurations["green"])
	}

	intersection.writeStatus()

	if elapsed := time.Now().Sub(intersection.start); elapsed < intersection.maxDuration {
		go intersection.switchLights()
	} else {
		// If the application has reached the max duration the outputStatus channel has to be closed.
		close(intersection.outputStatus)
	}
}
