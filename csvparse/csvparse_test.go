package csvparse_test

import (
	"github.com/amitkgupta/goodlearn/csvparse"
	"github.com/amitkgupta/goodlearn/csvparse/csvparseutilities"
	"github.com/amitkgupta/goodlearn/data/row"
	"github.com/amitkgupta/goodlearn/data/target"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Csvparse", func() {
	Describe("DatasetFromPath", func() {
		Context("Given a path to a file that doesn't exist", func() {
			It("Returns an error", func() {
				_, err := csvparse.DatasetFromPath("testassets/nonexistent.csv", 1, 3)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(csvparseutilities.UnableToOpenFileError{}))
			})
		})

		Context("Given a path to an empty file", func() {
			It("Returns an error", func() {
				_, err := csvparse.DatasetFromPath("testassets/empty.csv", 1, 3)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(csvparseutilities.UnableToReadTwoLinesError{}))

			})
		})

		Context("Given a path to a file with only one line (only headers)", func() {
			It("Returns an error", func() {
				_, err := csvparse.DatasetFromPath("testassets/oneline.csv", 1, 3)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(csvparseutilities.UnableToReadTwoLinesError{}))

			})
		})

		Context("Given a path to a file where lines have different number of comma-separated values", func() {
			It("Returns an error", func() {
				_, err := csvparse.DatasetFromPath("testassets/badrowwidths.csv", 1, 3)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(csvparseutilities.GenericError{}))

			})
		})

		Context("Given a path to a file with inconsistent types in the columns", func() {
			It("Returns an error", func() {
				_, err := csvparse.DatasetFromPath("testassets/badcolumntypes.csv", 1, 3)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(csvparseutilities.UnableToParseRowError{}))

			})
		})

		Context("Given a path to a file with the first data row having an unparseable float", func() {
			It("Returns an error", func() {
				_, err := csvparse.DatasetFromPath("testassets/badfloatfirstdatarow.csv", 1, 3)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(csvparseutilities.UnableToParseColumnTypesError{}))

			})
		})

		Context("Given a path to a file with a subsequent data row with an unparseable float", func() {
			It("Returns an error", func() {
				_, err := csvparse.DatasetFromPath("testassets/badfloatlaterdatarow.csv", 1, 3)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(csvparseutilities.UnableToParseRowError{}))

			})
		})

		Context("Given a mismatch in number of columns, target range start, and target range end", func() {
			It("Returns an error", func() {
				_, err := csvparse.DatasetFromPath("testassets/good.csv", 1, 6)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(csvparseutilities.UnableToCreateDatasetError{}))

			})
		})

		Context("Given a path to a good CSV file and target range", func() {
			It("Does not return an error", func() {
				_, err := csvparse.DatasetFromPath("testassets/good.csv", 1, 3)
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("Returns a good dataset", func() {
				dataset, _ := csvparse.DatasetFromPath("testassets/good.csv", 1, 3)

				Ω(dataset.AllFeaturesFloats()).Should(BeTrue())
				Ω(dataset.NumFeatures()).Should(Equal(3))
				Ω(dataset.NumRows()).Should(Equal(3))

				secondRow, err := dataset.Row(1)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(secondRow.Target().Equals(target.Target{"x", 3.2, "y"})).Should(BeTrue())

				floatFeatureSecondRow, ok := secondRow.(row.FloatFeatureRow)
				Ω(ok).Should(BeTrue())
				Ω(floatFeatureSecondRow.Features()).Should(Equal([]float64{-22e8, 1.0, 7.0}))
			})
		})
	})
})
