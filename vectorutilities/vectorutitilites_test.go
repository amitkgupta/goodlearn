package vectorutilities_test

import (
	"github.com/amitkgupta/goodlearn/vectorutilities"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Vector Utilities", func() {
	Describe("Add", func() {
		It("Adds", func() {
			x := []float64{1.01, 2, 3, 4}
			y := []float64{3.14, -1, 0, 0}

			Ω(vectorutilities.Add(x, y)).Should(Equal([]float64{4.15, 1, 3, 4}))
		})
	})

	Describe("Scale", func() {
		It("Scales", func() {
			x := []float64{1, 2, 3, 4}
			a := -2.33

			Ω(vectorutilities.Scale(a, x)).Should(Equal([]float64{-2.33, -4.66, -6.99, -9.32}))
		})
	})
})
