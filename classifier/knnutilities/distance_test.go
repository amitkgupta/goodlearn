package knnutilities_test

import (
	. "github.com/amitkgupta/goodlearn/classifier/knnutilities"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Euclidean", func() {
	x := []float64{1, 2, -3.14}
	y := []float64{-5, 6, 2.718}
	squareEuclideanDistance := 86.316164

	Context("When the square of the Euclidean distance is less than the bailout", func() {
		var bailout float64 = 90

		It("Returns the square distance", func() {
			Ω(Euclidean(x, y, bailout)).Should(BeNumerically("~", squareEuclideanDistance, 0.001))
		})
	})

	Context("When the square of the Euclidean distance is greater than or equal to the bailout", func() {
		var bailout float64 = 80

		It("Returns the bailout", func() {
			Ω(Euclidean(x, y, bailout)).Should(Equal(bailout))
		})
	})
})
