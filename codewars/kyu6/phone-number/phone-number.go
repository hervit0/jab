package phonenumber

import (
	"strconv"
	"log"
	"strings"
)

func CreatePhoneNumber(numbers [10]uint) string {
	numbersString := make([]string, 10)
	for i, _ := range numbersString {
	log.Println(strconv.FormatUint(numbers[i], 16))
		numbersString[i] = string(numbers[i])
	}

	log.Println(numbersString)
	log.Println(strings.Join(numbersString[0:2], ""))
	return ""
}
