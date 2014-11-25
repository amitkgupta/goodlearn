package gradientdescentestimator_test

import (
	"github.com/amitkgupta/goodlearn/data/columntype"
	"github.com/amitkgupta/goodlearn/data/dataset"
	gdeErrors "github.com/amitkgupta/goodlearn/errors/parameterestimator/gradientdescentestimatorerrors"
	"github.com/amitkgupta/goodlearn/parameterestimator"
	"github.com/amitkgupta/goodlearn/parameterestimator/gradientdescentestimator"

	"fmt"
	"math"
	"math/rand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gradient Descent Parameter Estimation", func() {
	var lossGradient gradientdescentestimator.ParameterizedLossGradient

	Describe("NewGradientDescentParameterEstimator", func() {
		Context("Given negative learning rate", func() {
			It("Returns an error", func() {
				_, err := gradientdescentestimator.NewGradientDescentParameterEstimator(
					-0.3,
					0.3,
					100,
					lossGradient,
				)

				Ω(err).Should(BeAssignableToTypeOf(gdeErrors.InvalidGDPEInitializationValuesError{}))
			})
		})

		Context("Given zero learning rate", func() {
			It("Returns an error", func() {
				_, err := gradientdescentestimator.NewGradientDescentParameterEstimator(
					0,
					0.3,
					100,
					lossGradient,
				)

				Ω(err).Should(BeAssignableToTypeOf(gdeErrors.InvalidGDPEInitializationValuesError{}))
			})
		})

		Context("Given negative precision", func() {
			It("Returns an error", func() {
				_, err := gradientdescentestimator.NewGradientDescentParameterEstimator(
					0.3,
					-0.3,
					100,
					lossGradient,
				)

				Ω(err).Should(BeAssignableToTypeOf(gdeErrors.InvalidGDPEInitializationValuesError{}))
			})
		})

		Context("Given zero precision", func() {
			It("Returns an error", func() {
				_, err := gradientdescentestimator.NewGradientDescentParameterEstimator(
					0.3,
					0,
					100,
					lossGradient,
				)

				Ω(err).Should(BeAssignableToTypeOf(gdeErrors.InvalidGDPEInitializationValuesError{}))
			})
		})

		Context("Given negative maximum number of iterations", func() {
			It("Returns an error", func() {
				_, err := gradientdescentestimator.NewGradientDescentParameterEstimator(
					0.3,
					0.3,
					-100,
					lossGradient,
				)

				Ω(err).Should(BeAssignableToTypeOf(gdeErrors.InvalidGDPEInitializationValuesError{}))
			})
		})

		Context("Given zero maximum number of iterations", func() {
			It("Returns an error", func() {
				_, err := gradientdescentestimator.NewGradientDescentParameterEstimator(
					0.3,
					0.3,
					0,
					lossGradient,
				)

				Ω(err).Should(BeAssignableToTypeOf(gdeErrors.InvalidGDPEInitializationValuesError{}))
			})
		})

		Context("Given positive learning rate, precisions, and max iterations", func() {
			It("Does not return an error", func() {
				_, err := gradientdescentestimator.NewGradientDescentParameterEstimator(
					0.3,
					0.3,
					100,
					lossGradient,
				)

				Ω(err).ShouldNot(HaveOccurred())
			})
		})
	})

	Describe("Train", func() {
		var trainingSet dataset.Dataset
		var estimator parameterestimator.ParameterEstimator

		BeforeEach(func() {
			var err error
			estimator, err = gradientdescentestimator.NewGradientDescentParameterEstimator(
				0.3,
				0.3,
				100,
				lossGradient,
			)
			Ω(err).ShouldNot(HaveOccurred())
		})

		Context("Given a dataset with non-float features", func() {
			BeforeEach(func() {
				columnTypes, err := columntype.StringsToColumnTypes([]string{"x", "1.0"})
				Ω(err).ShouldNot(HaveOccurred())

				trainingSet = dataset.NewDataset([]int{0}, []int{1}, columnTypes)

				err = trainingSet.AddRowFromStrings([]string{"hi", "24"})
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("Returns an error", func() {
				err := estimator.Train(trainingSet)
				Ω(err).Should(BeAssignableToTypeOf(gdeErrors.NonFloatFeaturesError{}))
			})
		})

		Context("Given a dataset with a non-float target", func() {
			BeforeEach(func() {
				columnTypes, err := columntype.StringsToColumnTypes([]string{"x", "1.0"})
				Ω(err).ShouldNot(HaveOccurred())

				trainingSet = dataset.NewDataset([]int{1}, []int{0}, columnTypes)

				err = trainingSet.AddRowFromStrings([]string{"hi", "24"})
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("Returns an error", func() {
				err := estimator.Train(trainingSet)
				Ω(err).Should(BeAssignableToTypeOf(gdeErrors.NonFloatTargetError{}))
			})
		})

		Context("Given a dataset with multiple target columns", func() {
			BeforeEach(func() {
				columnTypes, err := columntype.StringsToColumnTypes([]string{"1.0", "1.0", "1.0"})
				Ω(err).ShouldNot(HaveOccurred())

				trainingSet = dataset.NewDataset([]int{0}, []int{1, 2}, columnTypes)

				err = trainingSet.AddRowFromStrings([]string{"3.0", "24", "-1e8"})
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("Returns an error", func() {
				err := estimator.Train(trainingSet)
				Ω(err).Should(BeAssignableToTypeOf(gdeErrors.InvalidNumberOfTargetsError{}))
			})
		})

		Context("Given a dataset with no features", func() {
			BeforeEach(func() {
				columnTypes, err := columntype.StringsToColumnTypes([]string{"1.0"})
				Ω(err).ShouldNot(HaveOccurred())

				trainingSet = dataset.NewDataset([]int{}, []int{0}, columnTypes)

				err = trainingSet.AddRowFromStrings([]string{"3.14"})
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("Returns an error", func() {
				err := estimator.Train(trainingSet)
				Ω(err).Should(BeAssignableToTypeOf(gdeErrors.NoFeaturesError{}))
			})
		})

		Context("Given a valid dataset", func() {
			BeforeEach(func() {
				columnTypes, err := columntype.StringsToColumnTypes([]string{"1.0", "1.0"})
				Ω(err).ShouldNot(HaveOccurred())

				trainingSet = dataset.NewDataset([]int{0}, []int{1}, columnTypes)

				err = trainingSet.AddRowFromStrings([]string{"-3.14", "24"})
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("Does not return an error", func() {
				err := estimator.Train(trainingSet)
				Ω(err).ShouldNot(HaveOccurred())
			})
		})
	})

	Describe("Estimate", func() {
		var estimator parameterestimator.ParameterEstimator
		var trainingSet dataset.Dataset

		BeforeEach(func() {
			var err error
			estimator, err = gradientdescentestimator.NewGradientDescentParameterEstimator(
				0.000001,
				0.001,
				50,
				lossGradient,
			)
			Ω(err).ShouldNot(HaveOccurred())
		})

		Context("When the estimator hasn't been trained", func() {
			It("Returns an error", func() {
				_, err := estimator.Estimate([]float64{0.02, 0.1, 0.1, 0.1})
				Ω(err).Should(BeAssignableToTypeOf(gdeErrors.UntrainedEstimatorError{}))
			})
		})

		Context("When the parameter guess is empty", func() {
			BeforeEach(func() {
				err := estimator.Train(trainingSet)
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("Returns an error", func() {
				_, err := estimator.Estimate([]float64{})
				Ω(err).Should(BeAssignableToTypeOf(gdeErrors.EmptyInitialParametersError{}))
			})
		})

		Context("When the loss gradient returns an error (e.g. because of mis-shaped initial guess)", func() {
			BeforeEach(func() {
				err := estimator.Train(trainingSet)
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("Returns an error", func() {
				_, err := estimator.Estimate([]float64{0.01})
				Ω(err).Should(Equal(testError{1}))
			})
		})

		Context("When trained and given a valid guess", func() {
			BeforeEach(func() {
				err := estimator.Train(trainingSet)
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("Does not return an error", func() {
				_, err := estimator.Estimate([]float64{0.02, 0.1, 0.1, 0.1})
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("Returns an estimate of the parameters", func() {
				estimatedParameters, _ := estimator.Estimate([]float64{0.0196, 0.1004, 0.1004, 0.0996})
				Ω(estimatedParameters).Should(HaveLen(4))

				trueParameters := []float64{0.02, 0.1, 0.1, 0.1}
				for i := 0; i < 4; i++ {
					Ω(estimatedParameters[i]).Should(BeNumerically("~", trueParameters[i], 0.0005))
				}
			})
		})

		BeforeEach(func() {
			columnTypes, err := columntype.StringsToColumnTypes([]string{"1.0", "1.0"})
			Ω(err).ShouldNot(HaveOccurred())

			trainingSet = dataset.NewDataset([]int{0}, []int{1}, columnTypes)

			for x := -1.0; x <= 1.0; x = x + 0.05 {
				y := 0.02*math.Pow(x, 3) + 0.1*math.Pow(x, 2) + 0.1*x + 0.1 + (0.01 * rand.NormFloat64())

				err = trainingSet.AddRowFromStrings([]string{
					fmt.Sprintf("%.10f", x),
					fmt.Sprintf("%.10f", y),
				})
				Ω(err).ShouldNot(HaveOccurred())
			}
		})
	})

	BeforeEach(func() {
		// least squares cubic fit with penalty for cubic coefficient
		lossGradient = func(guess, x []float64, y float64) ([]float64, error) {
			if len(guess) != 4 {
				return nil, testError{1}
			}

			if len(x) != 1 {
				return nil, testError{2}
			}

			x0 := x[0]
			a := guess[0]
			b := guess[1]
			c := guess[2]
			d := guess[3]

			z := 2 * (y - a*math.Pow(x0, 3) - b*math.Pow(x0, 2) - c*x0 - d)

			result := make([]float64, 4)
			result[0] = z*math.Pow(x0, 3) + 2*a
			result[1] = z * math.Pow(x0, 2)
			result[2] = z * x0
			result[3] = z

			return result, nil
		}
	})
})

type testError struct {
	code int
}

func (e testError) Error() string {
	return "test error"
}
