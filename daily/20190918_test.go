package daily_test

import (
	"github.com/hervit0/jab/daily"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestSum(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sum")
}

var _ = Describe("Sum", func() {
	Context("When there is a match", func() {
		It("returns true", func() {
			//Given
			list := []int{10, 15, 3, 7}
			sum := 17

			// When
			result := daily.Sum(list, sum)

			// Then
			Expect(result).To(Equal(true))
		})
	})

	Context("When there is no match", func() {
		It("returns true", func() {
			//Given
			list := []int{10, 15, 3, 8}
			sum := 17

			// When
			result := daily.Sum(list, sum)

			// Then
			Expect(result).To(Equal(false))
		})
	})

	Context("When there is match - target 16", func() {
		It("returns true", func() {
			//Given
			list := []int{8, 10, 15, 8}
			sum := 16

			// When
			result := daily.Sum(list, sum)

			// Then
			Expect(result).To(Equal(true))
		})
	})

	Context("When there is no match - target 16", func() {
		It("returns true", func() {
			//Given
			list := []int{8, 10, 15, 3}
			sum := 16

			// When
			result := daily.Sum(list, sum)

			// Then
			Expect(result).To(Equal(false))
		})
	})
})
