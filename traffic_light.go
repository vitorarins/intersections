package main

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
	color Color
}

func NewTrafficLight(color Color) *TrafficLight {
	tl := &TrafficLight{color}
	return tl
}

func (tl *TrafficLight) switchColor() {
	tl.color = (tl.color + 1) % 3 // as we have 3 possible colors
}
