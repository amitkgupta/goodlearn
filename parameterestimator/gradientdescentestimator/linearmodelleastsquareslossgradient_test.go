package gradientdescentestimator_test

import (
	"github.com/amitkgupta/goodlearn/data/columntype"
	"github.com/amitkgupta/goodlearn/data/dataset"
	"github.com/amitkgupta/goodlearn/parameterestimator"
	"github.com/amitkgupta/goodlearn/parameterestimator/gradientdescentestimator"

	"fmt"
	"math/rand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Linear Model Least Squares Parameter Estimation", func() {
	var estimator parameterestimator.ParameterEstimator
	var trueParameters []float64

	Context("Given reasonable learning rate, precision, maximum number of iterations, and training set", func() {
		BeforeEach(func() {
			trueParameters = []float64{2, -3, 4}

			var err error
			estimator, err = gradientdescentestimator.NewGradientDescentParameterEstimator(
				0.001,
				0.000005,
				1000,
				gradientdescentestimator.LinearModelLeastSquaresLossGradient,
			)
			Ω(err).ShouldNot(HaveOccurred())

			columnTypes, err := columntype.StringsToColumnTypes([]string{"1.0", "1.0", "1.0"})
			Ω(err).ShouldNot(HaveOccurred())

			trainingSet := dataset.NewDataset([]int{0, 1}, []int{2}, columnTypes)
			for i := 0; i < 20; i++ {
				for j := 0; j < 20; j++ {
					x0 := -1.9 + 0.2*float64(i)
					x1 := -1.9 + 0.2*float64(j)
					y := trueParameters[0]*x0 + trueParameters[1]*x1 + trueParameters[2] + 0.1*rand.NormFloat64()

					err = trainingSet.AddRowFromStrings([]string{
						fmt.Sprintf("%.10f", x0),
						fmt.Sprintf("%.10f", x1),
						fmt.Sprintf("%.10f", y),
					})
					Ω(err).ShouldNot(HaveOccurred())
				}
			}

			err = estimator.Train(trainingSet)
			Ω(err).ShouldNot(HaveOccurred())
		})

		Context("Given a mis-shaped initial parameter guess", func() {
			It("Returns an error", func() {
				_, err := estimator.Estimate([]float64{1.75, -3.25})
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("Given a valid initial parameter guess", func() {
			It("Does not return an error", func() {
				_, err := estimator.Estimate([]float64{1.75, -3.25, 4.25})
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("Estimates the true parameters", func() {
				estimatedParameters, _ := estimator.Estimate([]float64{1.75, -3.25, 4.25})
				Ω(estimatedParameters).Should(HaveLen(3))

				for i := 0; i < 3; i++ {
					Ω(estimatedParameters[i]).Should(BeNumerically("~", trueParameters[i], 0.005))
				}
			})
		})
	})
})
