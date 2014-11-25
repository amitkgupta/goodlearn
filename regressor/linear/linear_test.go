package linear_test

import (
	"github.com/amitkgupta/goodlearn/data/columntype"
	"github.com/amitkgupta/goodlearn/data/dataset"
	"github.com/amitkgupta/goodlearn/data/row"
	"github.com/amitkgupta/goodlearn/data/slice"
	"github.com/amitkgupta/goodlearn/errors/regressor/linearerrors"
	"github.com/amitkgupta/goodlearn/regressor"
	"github.com/amitkgupta/goodlearn/regressor/linear"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LinearRegressor", func() {
	var linearRegressor regressor.Regressor

	Describe("Train", func() {
		var trainingData dataset.Dataset

		BeforeEach(func() {
			linearRegressor = linear.NewLinearRegressor()
		})

		Context("When the dataset's features are not all floats", func() {
			BeforeEach(func() {
				columnTypes, err := columntype.StringsToColumnTypes([]string{"3.3", "bye", "0"})
				Ω(err).ShouldNot(HaveOccurred())

				trainingData = dataset.NewDataset([]int{1, 2}, []int{0}, columnTypes)
			})

			It("Returns an error", func() {
				err := linearRegressor.Train(trainingData)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(linearerrors.NonFloatFeaturesTrainingSetError{}))
			})
		})

		Context("When the dataset's targets are not all floats", func() {
			BeforeEach(func() {
				columnTypes, err := columntype.StringsToColumnTypes([]string{"hi", "2.3", "0"})
				Ω(err).ShouldNot(HaveOccurred())

				trainingData = dataset.NewDataset([]int{1, 2}, []int{0}, columnTypes)
			})

			It("Returns an error", func() {
				err := linearRegressor.Train(trainingData)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(linearerrors.NonFloatTargetsTrainingSetError{}))
			})
		})

		Context("When the dataset has zero target values", func() {
			BeforeEach(func() {
				columnTypes, err := columntype.StringsToColumnTypes([]string{"1.3", "2.0"})
				Ω(err).ShouldNot(HaveOccurred())

				trainingData = dataset.NewDataset([]int{1, 0}, []int{}, columnTypes)
			})

			It("Returns an error", func() {
				err := linearRegressor.Train(trainingData)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(linearerrors.InvalidNumberOfTargetsError{}))
			})
		})

		Context("When the dataset has more than one target value", func() {
			BeforeEach(func() {
				columnTypes, err := columntype.StringsToColumnTypes([]string{"1.3", "2.2", "0"})
				Ω(err).ShouldNot(HaveOccurred())

				trainingData = dataset.NewDataset([]int{1}, []int{0, 2}, columnTypes)
			})

			It("Returns an error", func() {
				err := linearRegressor.Train(trainingData)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(linearerrors.InvalidNumberOfTargetsError{}))
			})
		})

		Context("When the dataset has no features", func() {
			BeforeEach(func() {
				columnTypes, err := columntype.StringsToColumnTypes([]string{"1.3"})
				Ω(err).ShouldNot(HaveOccurred())

				trainingData = dataset.NewDataset([]int{}, []int{0}, columnTypes)
			})

			It("Returns an error", func() {
				err := linearRegressor.Train(trainingData)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(linearerrors.NoFeaturesError{}))
			})
		})

		Context("When the dataset is valid", func() {
			BeforeEach(func() {
				columnTypes, err := columntype.StringsToColumnTypes([]string{"1.2", "3.4", "5.6"})
				Ω(err).ShouldNot(HaveOccurred())

				trainingData = dataset.NewDataset([]int{1, 2}, []int{0}, columnTypes)

				err = trainingData.AddRowFromStrings([]string{"1.2", "3.0", "0.5"})
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("Doesn't return an error", func() {
				err := linearRegressor.Train(trainingData)
				Ω(err).ShouldNot(HaveOccurred())
			})
		})
	})

	Describe("Predict", func() {
		var testRow row.Row
		var emptyTarget slice.Slice
		var columnTypes []columntype.ColumnType
		var err error

		BeforeEach(func() {
			linearRegressor = linear.NewLinearRegressor()

			columnTypes, err = columntype.StringsToColumnTypes([]string{"0", "0", "0"})
			Ω(err).ShouldNot(HaveOccurred())

			emptyTarget, err = slice.SliceFromRawValues(true, []int{}, columnTypes, []float64{})
			Ω(err).ShouldNot(HaveOccurred())
		})

		Context("When the regressor hasn't been trained", func() {
			BeforeEach(func() {
				features, err := slice.SliceFromRawValues(true, []int{1}, columnTypes, []float64{0, 1, 2})
				Ω(err).ShouldNot(HaveOccurred())

				testRow = row.NewRow(features, emptyTarget, 1)
			})

			It("Returns an error", func() {
				_, err := linearRegressor.Predict(testRow)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(linearerrors.UntrainedRegressorError{}))
			})
		})

		Context("When the regressor has been trained", func() {
			BeforeEach(func() {
				trainingData := dataset.NewDataset([]int{1, 2}, []int{0}, columnTypes)

				err = trainingData.AddRowFromStrings([]string{"-0.002001", "2", "3"})
				Ω(err).ShouldNot(HaveOccurred())

				err = trainingData.AddRowFromStrings([]string{"-0.001001", "2", "2"})
				Ω(err).ShouldNot(HaveOccurred())

				err = trainingData.AddRowFromStrings([]string{"-0.000999", "3", "3"})
				Ω(err).ShouldNot(HaveOccurred())

				err = linearRegressor.Train(trainingData)
				Ω(err).ShouldNot(HaveOccurred())
			})

			Context("When number of test features does not equal number of training features", func() {
				BeforeEach(func() {
					features, err := slice.SliceFromRawValues(true, []int{1}, columnTypes, []float64{0, 1, 2})
					Ω(err).ShouldNot(HaveOccurred())

					testRow = row.NewRow(features, emptyTarget, 1)
				})

				It("Returns an error", func() {
					_, err := linearRegressor.Predict(testRow)
					Ω(err).Should(HaveOccurred())
					Ω(err).Should(BeAssignableToTypeOf(linearerrors.RowLengthMismatchError{}))
				})
			})

			Context("When the test row's features are not all floats", func() {
				BeforeEach(func() {
					features, err := slice.SliceFromRawValues(false, []int{0, 1}, columnTypes, []float64{0, 1, 2})
					Ω(err).ShouldNot(HaveOccurred())

					testRow = row.NewRow(features, emptyTarget, 2)
				})

				It("Returns an error", func() {
					_, err := linearRegressor.Predict(testRow)
					Ω(err).Should(HaveOccurred())
					Ω(err).Should(BeAssignableToTypeOf(linearerrors.NonFloatFeaturesTestRowError{}))
				})
			})

			Context("When the test row is compatible with the training data", func() {
				BeforeEach(func() {
					features, err := slice.SliceFromRawValues(true, []int{1, 2}, columnTypes, []float64{0, 3.3, 1.0})
					Ω(err).ShouldNot(HaveOccurred())

					testRow = row.NewRow(features, emptyTarget, 2)
				})

				It("Predicts the target value for the test row", func() {
					predictedTarget, err := linearRegressor.Predict(testRow)
					Ω(err).ShouldNot(HaveOccurred())

					Ω(predictedTarget).Should(BeNumerically("~", 0.0013, 0.0001))
				})
			})
		})
	})
})
