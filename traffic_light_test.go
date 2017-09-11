package main

import (
	// "fmt"
	"testing"
	"time"
)

var testLightColorDurations = []time.Duration{
	(20 * time.Millisecond), // Green
	10 * time.Millisecond,   // Yellow
	30 * time.Millisecond,   // Red
}

var tl *TrafficLight = NewTrafficLight(Green, testLightColorDurations)

func TestSwitchToYellow(t *testing.T) {
	// fmt.Printf("My color before: %s\n", tl.color)

	// time.Sleep(testLightColorDurations[Green] - time.Millisecond)
	if tl.color != Yellow {
		t.Errorf("Expected '%s' but got '%s'", "Yellow", tl.color)
	}
}

func TestSwitchToRed(t *testing.T) {
	// fmt.Printf("My color before: %s\n", tl.color)
	time.Sleep(testLightColorDurations[Yellow] - time.Millisecond)
	if tl.color != Red {
		t.Errorf("Expected '%s' but got '%s'", "Red", tl.color)
	}
}

func TestSwitchToGreen(t *testing.T) {
	// fmt.Printf("My color before: %s\n", tl.color)
	time.Sleep(testLightColorDurations[Red] - time.Millisecond)
	if tl.color != Green {
		t.Errorf("Expected '%s' but got '%s'", "Green", tl.color)
	}
}
