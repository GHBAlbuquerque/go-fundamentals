package main

import "testing"

var tests = []struct {
	name     string
	dividen  float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 5.0, 20.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
}

// parametrized tests using struct
func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividen, tt.divisor)

		if tt.isErr {
			if err == nil {
				t.Error("Expected error but did not get")
			}
		} else {
			if err != nil {
				t.Error("Did not expect error but got one:", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("Expected %f but got %f", tt.expected, got)
		}
	}
}

// simple unit tests
func TestDivide(t *testing.T) {
	_, err := divide(100.0, 20.0)

	if err != nil {
		t.Error("Got an error when we should not have", err.Error())
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(100.0, 0.0)

	if err == nil {
		t.Error("Should have gotten an error")
	}
}
