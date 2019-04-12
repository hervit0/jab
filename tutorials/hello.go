// go run tutorials/hello.go
package tutorials

import (
	"fmt"
	"math/rand"
	"math"
)

const wow string = "oh wow"

func hello() {
	fmt.Println("Hello, 世界")
	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	var a, b string = "231", "lol"
	fmt.Println(a)
	fmt.Println(b)

	lol := "21"
	fmt.Println(lol)

	fmt.Println(wow)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))
}
