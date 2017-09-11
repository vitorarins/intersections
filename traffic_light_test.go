package main

import (
	"testing"
)

func TestSwitchColor(t *testing.T) {
	tl := NewTrafficLight(Green)

	if tl.color != Yellow {
		t.Errorf("Expected '%s' but got '%s'", "Yellow", tl.color)
	}
}
