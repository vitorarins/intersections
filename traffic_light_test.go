package main

import (
	"testing"
	"time"
)

var testLightColorDurations = []time.Duration{
	200 * time.Millisecond, // Green
	100 * time.Millisecond, // Yellow
	300 * time.Millisecond, // Red
}

var tl *TrafficLight = NewTrafficLight(Green, testLightColorDurations)

func TestSwitchToYellow(t *testing.T) {
	time.Sleep(100 * time.Millisecond)
	if tl.color != Yellow {
		t.Errorf("Expected '%s' but got '%s'", "Yellow", tl.color)
	}
}

func TestSwitchToRed(t *testing.T) {
	time.Sleep(testLightColorDurations[Yellow] - time.Millisecond)
	if tl.color != Red {
		t.Errorf("Expected '%s' but got '%s'", "Red", tl.color)
	}
}

func TestSwitchToGreen(t *testing.T) {
	time.Sleep(testLightColorDurations[Red] - time.Millisecond)
	if tl.color != Green {
		t.Errorf("Expected '%s' but got '%s'", "Green", tl.color)
	}
}
