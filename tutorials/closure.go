// go run tutorials/closure.go
package main

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
func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
