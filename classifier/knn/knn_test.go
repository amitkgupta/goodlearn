package knn_test

import (
	"github.com/amitkgupta/goodlearn/classifier"
	"github.com/amitkgupta/goodlearn/classifier/knn"
	"github.com/amitkgupta/goodlearn/data/columntype"
	"github.com/amitkgupta/goodlearn/data/dataset"
	"github.com/amitkgupta/goodlearn/data/dataset/inmemorydataset"
	"github.com/amitkgupta/goodlearn/data/row"
	"github.com/amitkgupta/goodlearn/data/target"

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

				trainingData, err = inmemorydataset.NewDataset(0, 0, columnTypes)
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("Returns an error", func() {
				err := kNNClassifier.Train(trainingData)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(knn.EmptyTrainingDatasetError{}))
			})
		})

		Context("When the dataset's features are not all floats", func() {
			BeforeEach(func() {
				columnTypes, err := columntype.StringsToColumnTypes([]string{"hi", "bye", "0"})
				Ω(err).ShouldNot(HaveOccurred())

				trainingData, err = inmemorydataset.NewDataset(0, 0, columnTypes)
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("Returns an error", func() {
				err := kNNClassifier.Train(trainingData)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(knn.NonFloatFeaturesTrainingSetError{}))
			})
		})

		Context("When the dataset is valid", func() {
			BeforeEach(func() {
				columnTypes, err := columntype.StringsToColumnTypes([]string{"hi", "0", "0"})
				Ω(err).ShouldNot(HaveOccurred())

				trainingData, err = inmemorydataset.NewDataset(0, 0, columnTypes)
				Ω(err).ShouldNot(HaveOccurred())

				err = trainingData.AddRowFromStrings(0, 0, columnTypes, []string{"hi", "0", "0"})
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("Doesn't return an error", func() {
				err := kNNClassifier.Train(trainingData)
				Ω(err).ShouldNot(HaveOccurred())
			})
		})
	})

	Describe("Classify", func() {
		var testRow *row.Row

		BeforeEach(func() {
			kNNClassifier, _ = knn.NewKNNClassifier(1)
		})

		Context("When the classifier hasn't been trained", func() {
			BeforeEach(func() {
				testRow = row.UnsafeNewRow(target.Target{"bye"}, []float64{1}, true)
			})

			It("Returns an error", func() {
				_, err := kNNClassifier.Classify(testRow)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(knn.UntrainedClassifierError{}))
			})
		})

		Context("When the classifier has been trained", func() {
			BeforeEach(func() {
				columnTypes, err := columntype.StringsToColumnTypes([]string{"hi", "0", "0"})
				Ω(err).ShouldNot(HaveOccurred())

				trainingData, err := inmemorydataset.NewDataset(0, 0, columnTypes)
				Ω(err).ShouldNot(HaveOccurred())

				err = trainingData.AddRowFromStrings(0, 0, columnTypes, []string{"hi", "0", "0"})
				Ω(err).ShouldNot(HaveOccurred())

				err = kNNClassifier.Train(trainingData)
				Ω(err).ShouldNot(HaveOccurred())
			})

			Context("When number of test features does not equal number of training features", func() {
				BeforeEach(func() {
					testRow = row.UnsafeNewRow(target.Target{}, []float64{1}, true)
				})

				It("Returns an error", func() {
					_, err := kNNClassifier.Classify(testRow)
					Ω(err).Should(HaveOccurred())
					Ω(err).Should(BeAssignableToTypeOf(knn.RowLengthMismatchError{}))
				})
			})

			Context("When the test row's features are not all floats", func() {
				BeforeEach(func() {
					testRow = row.UnsafeNewRow(target.Target{}, []float64{1, 2}, false)
				})

				It("Returns an error", func() {
					_, err := kNNClassifier.Classify(testRow)
					Ω(err).Should(HaveOccurred())
					Ω(err).Should(BeAssignableToTypeOf(knn.NonFloatFeaturesTestRowError{}))
				})
			})

			Context("When the test row is compatible with the training data", func() {
				BeforeEach(func() {
					testRow = row.UnsafeNewRow(target.Target{}, []float64{1, 2}, true)
				})

				It("Classifies the test row", func() {
					classifiedTarget, err := kNNClassifier.Classify(testRow)
					Ω(err).ShouldNot(HaveOccurred())
					Ω(classifiedTarget).Should(Equal(target.Target{"hi"}))
				})
			})
		})
	})
})
