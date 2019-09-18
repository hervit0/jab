package daily

//Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
//For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
//Bonus: Can you do this in one pass?

func Sum(numbers []int, sum int) bool {
	reverseNumbers := make([]int, len(numbers))

	for i, number := range numbers {
		reverseNumbers[i] = sum - number
	}

	var match bool
	for i, number := range numbers {
		partialReverseNumbers := append(reverseNumbers[:i], reverseNumbers[i+1:]...)
		for _, revNumber := range partialReverseNumbers {
			if revNumber == number {
				match = true
				break
			}
		}
	}

	return match
}
