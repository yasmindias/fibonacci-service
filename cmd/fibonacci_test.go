package cmd

import "testing"

func TestFibonacciRecursion(t *testing.T) {
	result := FibonacciRecursion(10)
	expected := 55

	if result != expected {
		t.Errorf("result %d expected %d", result, expected)
	}

}
