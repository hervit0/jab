package daily

//Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
//For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
//Bonus: Can you do this in one pass?

// Slightly faster, a 2N-ish algo as it needs to lookup into a map
func Sum(numbers []int, sum int) bool {
	dict := map[int]map[int]bool{}

	for i, number := range numbers {
		reverseNum := sum - number
		indexesDict, ok := dict[reverseNum]
		if !ok {
			indexesDict = map[int]bool{}
		}
		indexesDict[i] = true
		dict[reverseNum] = indexesDict
	}

	var match bool
	for i, number := range numbers {
		indexesDict, ok := dict[number]
		if ok {
			delete(indexesDict, i)
			if len(indexesDict) != 0 {
				match = true
				break
			}
		}
	}

	return match
}

// This is an N^2 algo
func SumSlow(numbers []int, sum int) bool {
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
