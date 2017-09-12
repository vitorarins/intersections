package main

// Color represents a light color
type Color int

// The possible colors for a traffic light
const (
	Green Color = iota
	Yellow
	Red
)

var colors = []string{"Green", "Yellow", "Red"}

func (c Color) String() string {
	return colors[c]
}

// TrafficLight represents one or a pair of virtual traffic lights
type TrafficLight struct {
	color Color
}

// NewTrafficLight allocates a new TrafficLight with color
func NewTrafficLight(color Color) *TrafficLight {
	tl := &TrafficLight{color}
	return tl
}

func (tl *TrafficLight) switchColor() {
	tl.color = (tl.color + 1) % 3 // as we have 3 possible colors
}
