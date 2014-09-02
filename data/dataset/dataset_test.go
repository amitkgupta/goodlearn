package dataset_test

import (
	"github.com/amitkgupta/goodlearn/data/columntype"
	"github.com/amitkgupta/goodlearn/data/dataset"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dataset", func() {
	Describe("NewDataset", func() {
		var columnTypes []columntype.ColumnType

		BeforeEach(func() {
			var err error
			columnTypes, err = columntype.StringsToColumnTypes([]string{"hi", "0", "0", "hi"})
			Ω(err).ShouldNot(HaveOccurred())
		})

		Context("When the target start is greater than target end", func() {
			It("Returns an error", func() {
				_, err := dataset.NewDataset(1, 0, columnTypes)
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("When the target start is negative", func() {
			It("Returns an error", func() {
				_, err := dataset.NewDataset(-1, 0, columnTypes)
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("When the target end goes beyond the columns", func() {
			It("Returns an error", func() {
				_, err := dataset.NewDataset(1, 4, columnTypes)
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("When the target range covers all the columns", func() {
			It("Returns an error", func() {
				_, err := dataset.NewDataset(0, 3, columnTypes)
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("When the target range is valid with respec to the column types", func() {
			It("Does not return an error", func() {
				_, err := dataset.NewDataset(0, 1, columnTypes)
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("Returns a dataset", func() {
				result, _ := dataset.NewDataset(0, 1, columnTypes)
				_, ok := result.(dataset.Dataset)
				Ω(ok).Should(BeTrue())
			})
		})
	})
})
