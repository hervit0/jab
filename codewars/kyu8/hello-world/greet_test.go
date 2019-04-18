// TODO: replace with your own tests (TDD). An example to get you started is included below.
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo></http:>
// Gomega Matcher Library <http://onsi.github.io/gomega></http:>

package helloworld

import (
  "testing"

	. "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

func TestGreet(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Add that block and the 'testing' framework in the import")
}

var _ = Describe("Example Tests", func() {
	It("should return string: 'hello world'", func() {
		Expect(greet()).To(Equal("hello world"))
	})
})
