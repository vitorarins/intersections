package main

import (
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

var colorDurations = []time.Duration{
	(4 * time.Minute) + (30 * time.Second),
	30 * time.Second,
	5 * time.Minute,
}

type TrafficLight struct {
	color Color
	timer *time.Timer
}

func NewTrafficLight(color Color) *TrafficLight {
	timer := time.NewTimer(colorDurations[color])
	light := &TrafficLight{color, timer}
	go func() {
		<-timer.C
		light.switchColor()
	}()
	return light
}

func (tl *TrafficLight) switchColor() {
	tl.color = (tl.color + 1) % 3 // as we have 3 possible colors
	tl.timer = time.AfterFunc(colorDurations[tl.color], tl.switchColor)
}
