package csvparse_test

import (
	"github.com/amitkgupta/goodlearn/csvparse"
	"github.com/amitkgupta/goodlearn/data/slice"
	"github.com/amitkgupta/goodlearn/errors/csvparseerrors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Csvparse", func() {
	Describe("DatasetFromPath", func() {
		Context("Given a path to a file that doesn't exist", func() {
			It("Returns an error", func() {
				_, err := csvparse.DatasetFromPath("testassets/nonexistent.csv", 1, 4)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(csvparseerrors.UnableToOpenFileError{}))
			})
		})

		Context("Given a path to an empty file", func() {
			It("Returns an error", func() {
				_, err := csvparse.DatasetFromPath("testassets/empty.csv", 1, 4)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(csvparseerrors.UnableToReadTwoLinesError{}))

			})
		})

		Context("Given a path to a file with only one line (only headers)", func() {
			It("Returns an error", func() {
				_, err := csvparse.DatasetFromPath("testassets/oneline.csv", 1, 4)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(csvparseerrors.UnableToReadTwoLinesError{}))

			})
		})

		Context("Given a path to a file where lines have different number of comma-separated values", func() {
			It("Returns an error", func() {
				_, err := csvparse.DatasetFromPath("testassets/badrowwidths.csv", 1, 4)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(csvparseerrors.GenericError{}))

			})
		})

		Context("Given a path to a file with inconsistent types in the columns", func() {
			It("Returns an error", func() {
				_, err := csvparse.DatasetFromPath("testassets/badcolumntypes.csv", 1, 4)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(csvparseerrors.UnableToParseRowError{}))

			})
		})

		Context("Given a path to a file with the first data row having an unparseable float", func() {
			It("Returns an error", func() {
				_, err := csvparse.DatasetFromPath("testassets/badfloatfirstdatarow.csv", 1, 4)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(csvparseerrors.UnableToParseColumnTypesError{}))

			})
		})

		Context("Given a path to a file with a subsequent data row with an unparseable float", func() {
			It("Returns an error", func() {
				_, err := csvparse.DatasetFromPath("testassets/badfloatlaterdatarow.csv", 1, 4)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(csvparseerrors.UnableToParseRowError{}))

			})
		})

		Context("Given a mismatch in number of columns, target range start, and target range end", func() {
			It("Returns an error", func() {
				_, err := csvparse.DatasetFromPath("testassets/good.csv", 1, 7)
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(BeAssignableToTypeOf(csvparseerrors.TargetOutOfBoundsError{}))

			})
		})

		Context("Given a path to a good CSV file and target range", func() {
			It("Does not return an error", func() {
				_, err := csvparse.DatasetFromPath("testassets/good.csv", 1, 4)
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("Returns a good dataset", func() {
				dataset, _ := csvparse.DatasetFromPath("testassets/good.csv", 1, 4)

				Ω(dataset.AllFeaturesFloats()).Should(BeTrue())
				Ω(dataset.NumFeatures()).Should(Equal(3))
				Ω(dataset.NumRows()).Should(Equal(3))

				secondRow, err := dataset.Row(1)
				Ω(err).ShouldNot(HaveOccurred())

				target, ok := secondRow.Target().(slice.MixedSlice)
				Ω(ok).Should(BeTrue())
				Ω(target.Values()).Should(Equal([]interface{}{"x", 3.2, "y"}))

				features, ok := secondRow.Features().(slice.FloatSlice)
				Ω(ok).Should(BeTrue())
				Ω(features.Values()).Should(Equal([]float64{-22e8, 1.0, 7.0}))
			})
		})
	})
})
