package numerical_string

import (
	"strconv"
	"strings"
)

//Opening a buffer and close it at the very end
// Do we need to close the buffer?
//The map can be a map[rune]int
func Numericals(input string) string {
	counter := make(map[int32]int)
	var buf strings.Builder

	for _, c := range input {
		counter[c]++
		buf.WriteString(strconv.Itoa(counter[c]))
	}

	return buf.String()
}

// Concatenate string for each iteration
func NumericalsSlow(input string) string {
	counter := make(map[int32]int)

	var output string
	for _, c := range input {
		count, _ := counter[c]
		newCount := count + 1

		counter[c] = newCount
		output += strconv.Itoa(newCount)
	}

	return output
}
