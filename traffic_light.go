package main

import (
	// "fmt"
	"time"
)

type Color int

const (
	Green Color = iota
	Yellow
	Red
)

var colors = []string{"Green", "Yellow", "Red"}

func (c Color) String() string {
	return colors[c]
}

type TrafficLight struct {
	color          Color
	colorDurations []time.Duration
	timer          *time.Timer
}

func NewTrafficLight(color Color, colorDurations []time.Duration) *TrafficLight {
	light := &TrafficLight{color, colorDurations, nil}
	light.timer = time.AfterFunc(colorDurations[light.color], light.switchColor)
	return light
}

func (tl *TrafficLight) switchColor() {
	// fmt.Printf("My color before: %s\n", tl.color)
	tl.color = (tl.color + 1) % 3 // as we have 3 possible colors
	// fmt.Printf("My color after: %s\n", tl.color)
	tl.timer = time.AfterFunc(tl.colorDurations[tl.color], tl.switchColor)
}
