// +build codewars
// TODO: replace with your own tests (TDD). An example to get you started is included below.
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo></http:>
// Gomega Matcher Library <http://onsi.github.io/gomega></http:>

package DNAtoRNA

import (
  "testing"

	. "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

func TestDNAtoRNA(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Add that block and the 'testing' framework in the import")
}

var _ = Describe("Test Example", func() {
  It("GCAT", func() {
    Expect(DNAtoRNA("GCAT")).To(Equal("GCAU"))
    Expect(DNAtoRNA("TTTT")).To(Equal("UUUU"))
  })
})
