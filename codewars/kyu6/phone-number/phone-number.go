package phonenumber

import (
	"log"
	"strconv"
	"strings"
)

func CreatePhoneNumber(numbers [10]uint) string {
	numbersString := make([]string, 10)
	for i, _ := range numbersString {
		numbersString[i] = string(strconv.Itoa(int(numbers[i])))
	}

	log.Println(numbersString)
	log.Println(strings.Join(numbersString[0:3], ""))
	return ""
}
