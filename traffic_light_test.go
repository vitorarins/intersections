package main

import (
	"testing"
)

var tl = NewTrafficLight(Green)

func TestSwitchToYellow(t *testing.T) {
	tl.switchColor()
	if tl.color != Yellow {
		t.Errorf("Expected '%s' but got '%s'", "Yellow", tl.color)
	}
}

func TestSwitchToRed(t *testing.T) {
	tl.switchColor()
	if tl.color != Red {
		t.Errorf("Expected '%s' but got '%s'", "Red", tl.color)
	}
}

func TestSwitchToGreen(t *testing.T) {
	tl.switchColor()
	if tl.color != Green {
		t.Errorf("Expected '%s' but got '%s'", "Green", tl.color)
	}
}
