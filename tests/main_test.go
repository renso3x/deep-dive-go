package main

import "testing"

func TestCalculate(t *testing.T) {
	if calculate(2) != 4 {
		t.Error("expected 2 + 2 equal 4")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{99999, 100001},
	}

	for _, test := range tests {
		output := calculate(test.input)
		if output != test.expected {
			t.Error("Test failed: {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
