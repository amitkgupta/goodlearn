package gradientdescent_test

import (
	"github.com/amitkgupta/goodlearn/optimizer/gradientdescent"

	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GradientDescent", func() {
	var goodGradient func([]float64) ([]float64, error)

	BeforeEach(func() {
		goodGradient = func(x []float64) ([]float64, error) {
			g := make([]float64, len(x))
			for i, xi := range x {
				g[i] = 2 * xi
			}
			return g, nil
		}
	})

	Context("When given an empty initial guess", func() {
		It("Returns an error", func() {
			_, err := gradientdescent.GradientDescent([]float64{}, 0.05, 0.0005, 100000, goodGradient)
			Ω(err).Should(HaveOccurred())
		})
	})

	Context("When the given gradient function returns an error", func() {
		It("Returns an error", func() {
			badGradient := func(x []float64) ([]float64, error) {
				return nil, errors.New("I'm bad")
			}
			_, err := gradientdescent.GradientDescent([]float64{0.3, -0.4}, 0.05, 0.0005, 100000, badGradient)
			Ω(err).Should(HaveOccurred())
		})
	})

	Context("When given reasonable inputs", func() {
		var estimatedArgMin []float64
		var err error

		BeforeEach(func() {
			estimatedArgMin, err = gradientdescent.GradientDescent([]float64{0.3, -0.4}, 0.05, 0.0005, 100000, goodGradient)
		})

		It("Does not return an error", func() {
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Returns a reasonable output (local arg-min of the function with given gradient)", func() {
			Ω(estimatedArgMin[0]).Should(BeNumerically("~", 0.0, 0.005))
			Ω(estimatedArgMin[1]).Should(BeNumerically("~", 0.0, 0.005))
		})
	})

	Context("Context given a relatively small maximum for number of iterations", func() {
		var estimatedArgMin []float64
		var err error

		BeforeEach(func() {
			estimatedArgMin, err = gradientdescent.GradientDescent([]float64{0.3, -0.4}, 0.0005, 0.00000005, 1, goodGradient)
		})

		It("Does not return an error", func() {
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Returns an estimate close to the initial guess because it didn't run many iterations", func() {
			Ω(estimatedArgMin[0]).Should(BeNumerically("~", 0.3, 0.05))
			Ω(estimatedArgMin[1]).Should(BeNumerically("~", -0.4, 0.05))
		})
	})

	Context("Context given a relatively large precision", func() {
		var estimatedArgMin []float64
		var err error

		BeforeEach(func() {
			estimatedArgMin, err = gradientdescent.GradientDescent([]float64{0.3, -0.4}, 0.0005, 1, 1000000000, goodGradient)
		})

		It("Does not return an error", func() {
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("Returns an estimate close to the initial guess because it's already precise enough", func() {
			Ω(estimatedArgMin[0]).Should(BeNumerically("~", 0.3, 0.05))
			Ω(estimatedArgMin[1]).Should(BeNumerically("~", -0.4, 0.05))
		})
	})
})
