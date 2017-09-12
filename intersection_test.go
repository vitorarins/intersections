package main

import (
	"testing"
	"time"
)

func compare(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Expected '%s' but got '%s'", expected, actual)
	}

}

// keep the intersection duration low to have fast tests
var testMaxDuration = 2 * time.Second

// again keeping the time durations for each color low to have fast tests
var testColorDurations = map[string]time.Duration{
	"yellow": 200 * time.Millisecond,
	"switch": 500 * time.Millisecond,
}

var intersection *Intersection = NewIntersection(testMaxDuration, testColorDurations)

func TestInitialization(t *testing.T) {
	expected := "(N, S): Green, (E, W): Red - 00:00.00"
	actual := <-intersection.outputStatus
	compare(t, expected, actual)
}

func TestYellow(t *testing.T) {
	expected := "(N, S): Yellow, (E, W): Red - 00:00.30"
	actual := <-intersection.outputStatus
	compare(t, expected, actual)
}

func TestSwitch(t *testing.T) {
	expected := "(N, S): Red, (E, W): Green - 00:00.50"
	actual := <-intersection.outputStatus
	compare(t, expected, actual)
}

func TestAfterMaxDuration(t *testing.T) {
	expected := "(N, S): Green, (E, W): Red - 00:02.00"
	var actual string
	for actual = range intersection.outputStatus {
	}
	compare(t, expected, actual)
}
