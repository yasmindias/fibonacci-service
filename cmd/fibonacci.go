package cmd

import "math/big"

func FibonacciRecursion(n int) *big.Int {
	t0 := big.NewInt(0)
	t1 := big.NewInt(1)
	return additiveSequence(n, t0, t1)
}

func additiveSequence(n int, t0, t1 *big.Int) *big.Int {
	if n == 0 {
		return t0
	}
	if n == 1 {
		return t1
	}

	sum := big.NewInt(0)
	sum.Add(sum, t0)
	sum.Add(sum, t1)
	return additiveSequence(n-1, t1, sum)
}
