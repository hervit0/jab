package tutorials

import (
	"fmt"
	"log"
)

func intSeq() func() int {
	log.Println("Inside main function")
	i := 0

	return func() int {
		i++
		return i
	}
}

func closure() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
