package daily

//Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.
//For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].
//Follow-up: what if you can't use division?

// 2N speed with division
func Prod(numbers []int) []int {
	prod := 1

	for _, n := range numbers {
		prod *= n
	}

	prodNumbers := make([]int, len(numbers))
	for i, n := range numbers {
		prodNumbers[i] = prod / n
	}
	return prodNumbers
}
