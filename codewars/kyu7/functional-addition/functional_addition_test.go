// +build codewars

package functional_addition

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDNAtoRNA(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Add that block and the 'testing' framework in the import")
}

var _ = It("sample test", func() {
	Expect(Add(1)(3)).To(Equal(4))
})
