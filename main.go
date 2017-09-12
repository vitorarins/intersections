package main

import (
	"fmt"
	"time"
)

var maxDuration = 30 * time.Minute

var colorDurations = map[string]time.Duration{
	"yellow": 30 * time.Second,
	"switch": 5 * time.Minute,
}

func main() {
	intersection := NewIntersection(maxDuration, colorDurations)
	for s := range intersection.outputStatus {
		fmt.Println(s)
	}
}
