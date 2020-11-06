package cmd

import (
	"math/big"
	"testing"
)

func TestFibonacciRecursion(t *testing.T) {
	result := FibonacciRecursion(10)
	expected := big.NewInt(55)

	if result.Cmp(expected) != 0 {
		t.Errorf("result %d expected %d", result, expected)
	}

}
