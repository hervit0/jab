package numerical_string

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Add that block and the 'testing' framework in the import")
}

func dotest(input string, exp string) {
	var ans = Numericals(input)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {
	It("handles edge cases", func() {
		dotest("Hello, World!", "1112111121311")
		dotest("Inconceivable!", "11112211111121")
		dotest("hello hello", "11121122342")
		dotest("Hello", "11121")
		dotest("aaaaaaaaaaaa", "123456789101112")

		longString := String(60000)
		fmt.Printf(longString)
		//dotest(longString, "123456789101112")
	})
})

// Thanks: https://www.calhoun.io/creating-random-strings-in-go/
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}
