package tortoise

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Add that block and the 'testing' framework in the import")
}

func dotest(v1, v2, g int, exp [3]int) {
	var ans = Race(v1, v2, g)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		dotest(720, 850, 70, [3]int{0, 32, 18})
		dotest(820, 81, 550, [3]int{-1, -1, -1})
		dotest(80, 91, 37, [3]int{3, 21, 49})
	})
})
