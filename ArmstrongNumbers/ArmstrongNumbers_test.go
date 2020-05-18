package main

import (
	"testing"
)

func TestExponent(t *testing.T) {
	var expect float64 = 27
	got := exponent(3, 3)
	if got != expect {
		t.Errorf("expected: %f, got: %f", expect, got)
	}
}

func TestGetDigit(t *testing.T) {
	expected := []int{3, 7, 1}
	got, _ := getDigits(371)
	for i, _ := range got {
		if got[i] != expected[i] {
			t.Errorf("expected: %d, got %d", expected, got)
		}
	}
}

func TestArmstrong(t *testing.T) {
	expected := true
	got, _ := isArmstrong(371)
	if expected != got {
		t.Errorf("expected: %t, got: %t", expected, got)
	}
}
