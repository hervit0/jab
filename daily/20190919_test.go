package daily_test

import (
	"github.com/hervit0/jab/daily"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestProd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Prod")
}

var _ = Describe("Prod", func() {
	Context("When there is a match", func() {
		It("returns true", func() {
			//Given
			list := []int{1, 2, 3, 4, 5}

			// When
			result := daily.Prod(list)

			// Then
			Expect(result).To(Equal([]int{120, 60, 40, 30, 24}))
		})
	})
})
