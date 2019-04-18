package helloworld

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Add that block and the 'testing' framework in the import")
}

func TestGreet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Add that block and the 'testing' framework in the import")
}

var _ = Describe("Example Tests", func() {
	It("should return string: 'hello world'", func() {
		Expect(greet()).To(Equal("hello world"))
	})
})
