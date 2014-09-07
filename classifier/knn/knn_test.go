package knn_test

import (
	"github.com/amitkgupta/goodlearn/classifier"
	"github.com/amitkgupta/goodlearn/classifier/knn"
	"github.com/amitkgupta/goodlearn/data/columntype"
	"github.com/amitkgupta/goodlearn/data/dataset"
	"github.com/amitkgupta/goodlearn/data/row"
	"github.com/amitkgupta/goodlearn/data/slice"
	"github.com/amitkgupta/goodlearn/errors/classifier/knnerrors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KNNClassifier", func() {
	var kNNClassifier classifier.Classifier

	Describe("NewKNNClassifier", func() {
		Context("When given 0 for k", func() {
			It("Returns an error", func() {
				_, err := knn.NewKNNClassifier(0)
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("When given a negative value for k", func() {
			It("Returns an error", func() {
				_, err := knn.NewKNNClassifier(-3)
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("When given a positive value for k", func() {
			It("Returns an error", func() {
				_, err := knn.NewKNNClassifier(5)
				Ω(err).ShouldNot(HaveOccurred())
			})
		})
	})

	Describe("Train", func() {
		var trainingData dataset.Dataset

		BeforeEach(func() {
			kNNClassifier, _ = knn.NewKNNClassifier(1)
		})

		Context("When the dataset is empty", func() {
			BeforeEach(func() {
				columnTypes, err := columntype.StringsToColumnTypes([]string{"hi", "0", "0"})
				Ω(err).ShouldNot(HaveOccurred())

				trainingData = dataset.NewDataset([]int{1, 2}, []int{0}, columnTypes)
			})

			It("Returns an error", func() {
				err := kNNClassifier.Train(trainingData)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(knnerrors.EmptyTrainingDatasetError{}))
			})
		})

		Context("When the dataset's features are not all floats", func() {
			BeforeEach(func() {
				columnTypes, err := columntype.StringsToColumnTypes([]string{"hi", "bye", "0"})
				Ω(err).ShouldNot(HaveOccurred())

				trainingData = dataset.NewDataset([]int{1, 2}, []int{0}, columnTypes)
			})

			It("Returns an error", func() {
				err := kNNClassifier.Train(trainingData)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(knnerrors.NonFloatFeaturesTrainingSetError{}))
			})
		})

		Context("When the dataset is valid", func() {
			BeforeEach(func() {
				columnTypes, err := columntype.StringsToColumnTypes([]string{"hi", "0", "0"})
				Ω(err).ShouldNot(HaveOccurred())

				trainingData = dataset.NewDataset([]int{1, 2}, []int{0}, columnTypes)

				err = trainingData.AddRowFromStrings([]string{"hi", "0", "0"})
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("Doesn't return an error", func() {
				err := kNNClassifier.Train(trainingData)
				Ω(err).ShouldNot(HaveOccurred())
			})
		})
	})

	Describe("Classify", func() {
		var testRow row.Row
		var emptyTarget slice.Slice
		var columnTypes []columntype.ColumnType
		var err error
		var helloRaw float64

		BeforeEach(func() {
			kNNClassifier, _ = knn.NewKNNClassifier(1)

			columnTypes, err = columntype.StringsToColumnTypes([]string{"hi", "0", "0"})
			Ω(err).ShouldNot(HaveOccurred())

			emptyTarget, err = slice.SliceFromRawValues(true, []int{}, columnTypes, []float64{})
			Ω(err).ShouldNot(HaveOccurred())

			columnTypes, err := columntype.StringsToColumnTypes([]string{"hi", "0", "0"})
			Ω(err).ShouldNot(HaveOccurred())

			helloRaw, err = columnTypes[0].PersistRawFromString("val")
			Ω(err).ShouldNot(HaveOccurred())
		})

		Context("When the classifier hasn't been trained", func() {
			BeforeEach(func() {
				features, err := slice.SliceFromRawValues(true, []int{1}, columnTypes, []float64{helloRaw, 1, 2})
				Ω(err).ShouldNot(HaveOccurred())

				testRow = row.NewRow(features, emptyTarget, 1)
			})

			It("Returns an error", func() {
				_, err := kNNClassifier.Classify(testRow)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(knnerrors.UntrainedClassifierError{}))
			})
		})

		Context("When the classifier has been trained", func() {
			BeforeEach(func() {
				columnTypes, err = columntype.StringsToColumnTypes([]string{"hi", "0", "0"})
				Ω(err).ShouldNot(HaveOccurred())

				trainingData := dataset.NewDataset([]int{1, 2}, []int{0}, columnTypes)

				err = trainingData.AddRowFromStrings([]string{"hi", "0", "0"})
				Ω(err).ShouldNot(HaveOccurred())

				err = kNNClassifier.Train(trainingData)
				Ω(err).ShouldNot(HaveOccurred())
			})

			Context("When number of test features does not equal number of training features", func() {
				BeforeEach(func() {
					features, err := slice.SliceFromRawValues(true, []int{1}, columnTypes, []float64{helloRaw, 1, 2})
					Ω(err).ShouldNot(HaveOccurred())

					testRow = row.NewRow(features, emptyTarget, 1)
				})

				It("Returns an error", func() {
					_, err := kNNClassifier.Classify(testRow)
					Ω(err).Should(HaveOccurred())
					Ω(err).Should(BeAssignableToTypeOf(knnerrors.RowLengthMismatchError{}))
				})
			})

			Context("When the test row's features are not all floats", func() {
				BeforeEach(func() {
					features, err := slice.SliceFromRawValues(false, []int{0, 1}, columnTypes, []float64{helloRaw, 1, 2})
					Ω(err).ShouldNot(HaveOccurred())

					testRow = row.NewRow(features, emptyTarget, 2)
				})

				It("Returns an error", func() {
					_, err := kNNClassifier.Classify(testRow)
					Ω(err).Should(HaveOccurred())
					Ω(err).Should(BeAssignableToTypeOf(knnerrors.NonFloatFeaturesTestRowError{}))
				})
			})

			Context("When the test row is compatible with the training data", func() {
				BeforeEach(func() {
					features, err := slice.SliceFromRawValues(true, []int{1, 2}, columnTypes, []float64{helloRaw, 3.3, 1.0})
					Ω(err).ShouldNot(HaveOccurred())

					testRow = row.NewRow(features, emptyTarget, 2)
				})

				It("Classifies the test row", func() {
					classifiedTarget, err := kNNClassifier.Classify(testRow)
					Ω(err).ShouldNot(HaveOccurred())

					expectedTarget, err := slice.SliceFromRawValues(false, []int{0}, columnTypes, []float64{helloRaw, 99, 55})
					Ω(err).ShouldNot(HaveOccurred())
					Ω(classifiedTarget.Equals(expectedTarget)).Should(BeTrue())
				})
			})
		})
	})
})
