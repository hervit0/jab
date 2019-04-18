package DNAtoRNA

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Add that block and the 'testing' framework in the import")
}

var _ = Describe("Test Example", func() {
	It("GCAT", func() {
		Expect(DNAtoRNA("GCAT")).To(Equal("GCAU"))
		Expect(DNAtoRNA("TTTT")).To(Equal("UUUU"))
	})
})
