package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Fatalf("Add(2,3)=%d; expected %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(9, 8)
	expected := 1

	if result != expected {
		t.Fatalf("Subtract(9,8)=%d; expected %d,", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(4, 3)
	expected := 12

	if result != expected {
		t.Fatalf("Multiply(4,3)=%d; expected %d,", result, expected)
	}
}
