package sum_cubes

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Add that block and the 'testing' framework in the import")
}

func dotest(input int, exp int) {
	var ans = SumCubes(input)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {
	It("handles edge cases", func() {
		dotest(1, 1)
		dotest(2, 9)
		dotest(3, 36)
		dotest(4, 100)
		dotest(10, 3025)
		dotest(123, 58155876)

	})
})
