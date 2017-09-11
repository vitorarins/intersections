package main

import (
	"testing"
	"time"
)

func showError(t *testing.T, expected, actual string) {
	t.Errorf("Expected '%s' but got '%s'", expected, actual)
}

var testColorDurations = []time.Duration{
	3 * time.Second, // Green
	2 * time.Second, // Yellow
	5 * time.Second, // Red
}

var status chan string = Intersection(8*time.Second, testColorDurations)

func TestInitialization(t *testing.T) {
	expected := "(N, S): Green, (E, W): Red - 00:01"
	actual := <-status
	if actual != expected {
		showError(t, expected, actual)
	}
}

func TestYellow(t *testing.T) {
	expected := "(N, S): Yellow, (E, W): Red - 00:04"
	actual := <-status
	actual = <-status
	actual = <-status
	if actual != expected {
		showError(t, expected, actual)
	}
}

func TestSwitch(t *testing.T) {
	expected := "(N, S): Red, (E, W): Green - 00:06"
	actual := <-status
	actual = <-status
	if actual != expected {
		showError(t, expected, actual)
	}
}

func TestAfterDuration(t *testing.T) {
	expectedLastStatus := "(N, S): Red, (E, W): Green - 00:08"
	var actualLastStatus string
	for actualLastStatus = range status {
	}
	if actualLastStatus != expectedLastStatus {
		showError(t, expectedLastStatus, actualLastStatus)
	}
}
