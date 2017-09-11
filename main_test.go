package main

import (
	"testing"
	"time"
)

func showError(t *testing.T, expected, actual string) {
	t.Errorf("Expected '%s' but got '%s'", expected, actual)
}

var testColorDurations = []time.Duration{
	(2 * time.Second), // Green
	1 * time.Second,   // Yellow
	3 * time.Second,   // Red
}

var status chan string = Intersection(5*time.Second, testColorDurations)

func TestInitialization(t *testing.T) {
	expected := "NS: Green, EW: Red - 00:01"
	actual := <-status
	if actual != expected {
		showError(t, expected, actual)
	}
}

func TestYellow(t *testing.T) {
	expected := "NS: Yellow, EW: Red - 00:02"
	actual := <-status
	if actual != expected {
		showError(t, expected, actual)
	}
}

func TestSwitch(t *testing.T) {
	expected := "NS: Red, EW: Green - 00:03"
	actual := <-status
	if actual != expected {
		showError(t, expected, actual)
	}
}

func TestAfterDuration(t *testing.T) {
	expectedLastStatus := "NS: Red, EW: Green - 00:05"
	var actualLastStatus string
	for actualLastStatus = range status {
	}
	if actualLastStatus != expectedLastStatus {
		showError(t, expectedLastStatus, actualLastStatus)
	}
}
