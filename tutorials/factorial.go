package tutorials

// import "fmt"

// func fact(n int) int {
//     if n == 0 {
//         return 1
//     }
//     return n * fact(n-1)
// }

func fact(n int) int {
	return doFact(n, 1)
}

func doFact(n int, result int) int {
	if n == 0 {
		return result
	}
	return doFact(n-1, n*result)
}

// func main() {
// 	fmt.Println(fact(7))
// }
