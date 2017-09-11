package main

import (
	"testing"
	"time"
)

func showError(t *testing.T, expected, actual string) {
	t.Errorf("Expected '%s' but got '%s'", expected, actual)
}

var status chan string = Intersection(30)

func TestInitialization(t *testing.T) {
	tSub = func(n, s time.Time) time.Duration {
		return time.Duration(0)
	}
	expected := "NS: Green, EW: Red - 00:00"
	actual := <-status
	if actual != expected {
		showError(t, expected, actual)
	}
}

func TestYellow(t *testing.T) {
	tSub = func(n, s time.Time) time.Duration {
		return time.Duration((4 * time.Minute) + (30 * time.Second))
	}
	expected := "NS: Yellow, EW: Red - 04:30"
	actual := <-status
	if actual != expected {
		showError(t, expected, actual)
	}
}

func TestSwitch(t *testing.T) {
	tSub = func(n, s time.Time) time.Duration {
		return time.Duration(5 * time.Minute)
	}
	expected := "NS: Red, EW: Green - 05:00"
	actual := <-status
	if actual != expected {
		showError(t, expected, actual)
	}
}

func TestAfterDuration(t *testing.T) {
	tSub = func(n, s time.Time) time.Duration {
		return time.Duration(30 * time.Minute)
	}
	expected := ""
	actual := <-status
	if actual != expected {
		showError(t, expected, actual)
	}
}
