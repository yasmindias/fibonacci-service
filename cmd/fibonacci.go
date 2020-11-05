package cmd

func FibonacciRecursion(n int) int {
	return additiveSequence(n, 0, 1)
}

func additiveSequence(n, t0, t1 int) int {
	if n == 0 {
		return t0
	}
	if n == 1 {
		return t1
	}
	return additiveSequence(n-1, t1, t0+t1)
}
