package crossvalidation_test

import (
	"math/rand"
	"strconv"

	"github.com/amitkgupta/goodlearn/data/columntype"
	"github.com/amitkgupta/goodlearn/data/dataset"
	"github.com/amitkgupta/goodlearn/data/slice"
	"github.com/amitkgupta/goodlearn/evaluation/crossvalidation"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CrossValidation", func() {
	Describe("SplitDataset", func() {
		var (
			originalSet   dataset.Dataset
			trainingRatio float64

			trainingSet dataset.Dataset
			testSet     dataset.Dataset
			err         error
		)

		JustBeforeEach(func() {
			trainingSet, testSet, err = crossvalidation.SplitDataset(
				originalSet,
				trainingRatio,
				rand.NewSource(5330), // SEED
			)
		})

		BeforeEach(func() {
			columnTypes, columnTypesError := columntype.StringsToColumnTypes([]string{"0"})
			Ω(columnTypesError).ShouldNot(HaveOccurred())

			originalSet = dataset.NewDataset([]int{}, []int{0}, columnTypes)
		})

		Context("when the training ratio negative", func() {
			BeforeEach(func() {
				trainingRatio = -0.67
			})

			It("errors", func() {
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("when the training ratio is greater than 1", func() {
			BeforeEach(func() {
				trainingRatio = 1.67
			})

			It("errors", func() {
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("when the training ratio is valid", func() {
			BeforeEach(func() {
				trainingRatio = 0.80
			})

			Context("when the dataset has no rows", func() {
				It("errors", func() {
					Ω(err).Should(HaveOccurred())
				})
			})

			Context("when the dataset has rows and the ratio is valid", func() {
				var expectedTrainingTargets, expectedTestTargets []slice.Slice

				BeforeEach(func() {
					for i := 0; i < 20; i++ {
						originalSet.AddRowFromStrings([]string{strconv.Itoa(i)})
					}

					expectedTrainingTargets = makeSingleFloatTargets(16, 13, 1, 2, 19, 6, 5, 14, 10, 11, 4, 15, 8, 3, 9)
					expectedTestTargets = makeSingleFloatTargets(12, 7, 17, 0, 18)
				})

				It("splits the dataset into two according to the given ratio", func() {
					Ω(err).ShouldNot(HaveOccurred())

					Ω(trainingSet.NumRows()).Should(Equal(len(expectedTrainingTargets)))
					for i := 0; i < trainingSet.NumRows(); i++ {
						trainingRow, rowErr := trainingSet.Row(i)
						Ω(rowErr).ShouldNot(HaveOccurred())
						Ω(trainingRow.Target().Equals(expectedTrainingTargets[i])).Should(BeTrue())
					}

					Ω(testSet.NumRows()).Should(Equal(len(expectedTestTargets)))
					for i := 0; i < testSet.NumRows(); i++ {
						testRow, rowErr := testSet.Row(i)
						Ω(rowErr).ShouldNot(HaveOccurred())
						Ω(testRow.Target().Equals(expectedTestTargets[i])).Should(BeTrue())
					}
				})
			})
		})
	})
})

func makeSingleFloatTargets(floats ...float64) []slice.Slice {
	columnTypes, columnTypesError := columntype.StringsToColumnTypes([]string{"0"})
	Ω(columnTypesError).ShouldNot(HaveOccurred())

	targets := make([]slice.Slice, len(floats))
	for i, f := range floats {
		target, err := slice.SliceFromRawValues(true, []int{0}, columnTypes, []float64{f})
		Ω(err).ShouldNot(HaveOccurred())
		targets[i] = target
	}
	return targets
}
