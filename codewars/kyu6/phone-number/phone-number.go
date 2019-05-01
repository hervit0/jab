package phonenumber

import (
	"fmt"
	"strconv"
	"strings"
)

func CreatePhoneNumber(numbers [10]uint) string {
	numbersString := make([]string, 10)
	for i := range numbersString {
		numbersString[i] = string(strconv.Itoa(int(numbers[i])))
	}

	return fmt.Sprintf("(%s) %s-%s", join(numbersString[0:3]), join(numbersString[3:6]), join(numbersString[6:10]))
}

func join(portion []string) string {
	return strings.Join(portion, "")
}

// Lol
// func CreatePhoneNumber(n [10]uint) string {
//   return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", n[0], n[1], n[2], n[3], n[4], n[5], n[6], n[7], n[8], n[9])
// }
