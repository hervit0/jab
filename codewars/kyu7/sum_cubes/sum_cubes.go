package sum_cubes

import "math"

// Clean functional
func SumCubes(n int) int {
	return doSumCubes(n, 1)
}

func doSumCubes(n int, sum int) int {
	if n == 1 {
		return sum
	}
	return doSumCubes(n-1, int(math.Pow(float64(n), 3))+sum)
}

// Dirty functional - heap problem for non-optimized languages + it's not tail-recursive
func SumCubesDirty(n int) int {
	if n == 1 {
		return 1
	}
	return int(math.Pow(float64(n), 3)) + SumCubesDirty(n-1)
}
