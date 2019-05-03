package tortoise

import (
	// "log"
	"math"
)

func Race(v1, v2, g int) [3]int {
	if v1 >= v2 {
		return [3]int{-1, -1, -1}
	}

	timeInSecond := int(math.Trunc((float64(g) / float64(v2-v1) * 3600)))
	// log.Printf("test %v", math.Trunc((float64(g) / float64(v2-v1) * 3600)))
	// log.Printf("test %v", math.Ceil(((float64(g) / float64(v2-v1) * 3600))))
	hours, remainder := divmod(timeInSecond, 3600)
	minutes, seconds := divmod(remainder, 60)

	return [3]int{hours, minutes, seconds}
}

func divmod(numerator, denominator int) (int, int) {
	return numerator / denominator, numerator % denominator
}

// s := g * 3600 / (v2 - v1)
// return [3]int{s / 3600, (s / 60) % 60, s % 60}

// func divmod(numerator, denominator int) (quotient, remainder int) {
// 	quotient = numerator / denominator
// 	remainder = numerator % denominator
// 	return
// }
