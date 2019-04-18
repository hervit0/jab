package consecutive

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Add that block and the 'testing' framework in the import")
}

func dotest(strarr []string, k int, exp string) {
	var ans = LongestConsec(strarr, k)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {
	It("handles edge cases", func() {
		dotest([]string{}, 7, "")
		dotest([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 0, "")
		dotest([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 7, "")
	})

	It("should handle basic cases", func() {
		dotest([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2, "abigailtheta")
		dotest([]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, 1,
			"oocccffuucccjjjkkkjyyyeehh")
		dotest([]string{}, 3, "")
		dotest([]string{"itvayloxrp", "wkppqsztdkmvcuwvereiupccauycnjutlv", "vweqilsfytihvrzlaodfixoyxvyuyvgpck"}, 2,
			"wkppqsztdkmvcuwvereiupccauycnjutlvvweqilsfytihvrzlaodfixoyxvyuyvgpck")
	})
})
