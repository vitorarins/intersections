package main

import (
	"fmt"
	"time"
)

var maxDuration = 30 * time.Minute

var colorDurations = map[string]time.Duration{
	// this is how much time a light remains on yellow before making the switch from green to red.
	"yellow": 30 * time.Second,
	// this is how much time until a light changes from green to red or vice versa.
	"switch": 5 * time.Minute,
}

func main() {
	intersection := NewIntersection(maxDuration, colorDurations)
	for s := range intersection.outputStatus {
		fmt.Println(s)
	}
}
