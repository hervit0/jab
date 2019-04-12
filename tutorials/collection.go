// go run tutorials/collection.go
package main

import (
	"fmt"
)

func main() {
	// Array
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	undefinedLengthArray := [...]string{"oh", "wow"}
	fmt.Println("unknown length: ", undefinedLengthArray)

	// Slices
	fmt.Println("\n\nSLICES")
	t := []string{"g", "h", "i"}
	c := make([]string, len(t))
	copy(c, t)
	c = append(c, "l")
	fmt.Println("copied: ", c)
	fmt.Println("original: ", t)

	// Map
	fmt.Println("\n\nMAPS")
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	delete(m, "k2")
	fmt.Println("map:", m)

	v, prs := m["k2"]
	fmt.Println("prs:", prs)
	fmt.Println("v:", v)

	v1, prs1 := m["k1"]
	fmt.Println("prs:", prs1)
	fmt.Println("value:", v1)

	// Range
	fmt.Println("\n\nRANGE")
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
			sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
			if num == 3 {
					fmt.Println("index:", i)
			}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
			fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
			fmt.Println("key:", k)
	}

	for i, c := range "go" {
			fmt.Println(i, c)
	}
}
