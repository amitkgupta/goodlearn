package dataset_test

import (
	"github.com/amitkgupta/goodlearn/data/columntype"
	"github.com/amitkgupta/goodlearn/data/dataset"
	"github.com/amitkgupta/goodlearn/data/row"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dataset", func() {
	var ds dataset.Dataset

	Describe("AllFeaturesFloats", func() {
		Context("When all features are floats", func() {
			BeforeEach(func() {
				columnTypes, err := columntype.StringsToColumnTypes([]string{"1.0", "1.0"})
				Ω(err).ShouldNot(HaveOccurred())

				ds = dataset.NewDataset([]int{0, 1}, []int{}, columnTypes)
			})

			It("Returns true", func() {
				Ω(ds.AllFeaturesFloats()).Should(BeTrue())
			})
		})

		Context("When not all features are floats", func() {
			BeforeEach(func() {
				columnTypes, err := columntype.StringsToColumnTypes([]string{"x", "1.0"})
				Ω(err).ShouldNot(HaveOccurred())

				ds = dataset.NewDataset([]int{0, 1}, []int{}, columnTypes)
			})

			It("Returns true", func() {
				Ω(ds.AllFeaturesFloats()).Should(BeFalse())
			})
		})
	})

	Describe("AllTargetsFloats", func() {
		Context("When all targets are floats", func() {
			BeforeEach(func() {
				columnTypes, err := columntype.StringsToColumnTypes([]string{"1.0", "1.0"})
				Ω(err).ShouldNot(HaveOccurred())

				ds = dataset.NewDataset([]int{}, []int{0, 1}, columnTypes)
			})

			It("Returns true", func() {
				Ω(ds.AllTargetsFloats()).Should(BeTrue())
			})
		})

		Context("When not all targets are floats", func() {
			BeforeEach(func() {
				columnTypes, err := columntype.StringsToColumnTypes([]string{"x", "1.0"})
				Ω(err).ShouldNot(HaveOccurred())

				ds = dataset.NewDataset([]int{}, []int{0, 1}, columnTypes)
			})

			It("Returns true", func() {
				Ω(ds.AllTargetsFloats()).Should(BeFalse())
			})
		})
	})

	Describe("NumFeatures and NumTargets", func() {
		BeforeEach(func() {
			columnTypes, err := columntype.StringsToColumnTypes([]string{"1.0", "x", "x", "1.0", "x"})
			Ω(err).ShouldNot(HaveOccurred())

			ds = dataset.NewDataset([]int{1, 3}, []int{0, 2, 4}, columnTypes)
		})

		It("Returns the correct number of features and targets", func() {
			Ω(ds.NumFeatures()).Should(Equal(2))
			Ω(ds.NumTargets()).Should(Equal(3))
		})
	})

	Describe("Adding, Counting, and Getting rows", func() {
		BeforeEach(func() {
			columnTypes, err := columntype.StringsToColumnTypes([]string{"1.0", "x", "x", "1.0", "x"})
			Ω(err).ShouldNot(HaveOccurred())

			ds = dataset.NewDataset([]int{1, 3}, []int{0, 2, 4}, columnTypes)
		})

		Context("When the dataset is empty", func() {
			It("Has 0 rows", func() {
				Ω(ds.NumRows()).To(BeZero())
			})

			Context("When getting a row", func() {
				It("Returns an error", func() {
					_, err := ds.Row(0)
					Ω(err).Should(HaveOccurred())

					_, err = ds.Row(-1)
					Ω(err).Should(HaveOccurred())

					_, err = ds.Row(1)
					Ω(err).Should(HaveOccurred())
				})
			})

			Context("When adding a row", func() {
				var newRow row.Row
				var err error

				Context("When the row's length is incorrect", func() {
					BeforeEach(func() {
						err = ds.AddRowFromStrings([]string{"0.0", "hi", "mom", "94"})
					})

					It("Returns an error", func() {
						Ω(err).Should(HaveOccurred())
					})

					It("Has 0 rows", func() {
						Ω(ds.NumRows()).To(BeZero())
					})

					Context("When getting a row", func() {
						It("Returns an error", func() {
							_, err = ds.Row(0)
							Ω(err).Should(HaveOccurred())

							_, err = ds.Row(-1)
							Ω(err).Should(HaveOccurred())

							_, err = ds.Row(1)
							Ω(err).Should(HaveOccurred())
						})
					})
				})

				Context("When the row values don't match the column types", func() {
					BeforeEach(func() {
						err = ds.AddRowFromStrings([]string{"0.0", "hi", "mom", "shouldBeFloat", "word"})
					})

					It("Returns an error", func() {
						Ω(err).Should(HaveOccurred())
					})

					It("Has 0 rows", func() {
						Ω(ds.NumRows()).To(BeZero())
					})

					Context("When getting a row", func() {
						It("Returns an error", func() {
							_, err = ds.Row(0)
							Ω(err).Should(HaveOccurred())

							_, err = ds.Row(-1)
							Ω(err).Should(HaveOccurred())

							_, err = ds.Row(1)
							Ω(err).Should(HaveOccurred())
						})
					})
				})

				Context("When the row is valid", func() {
					BeforeEach(func() {
						err = ds.AddRowFromStrings([]string{"0.0", "hi", "mom", "94", "word"})
					})

					It("Does not return an error", func() {
						Ω(err).ShouldNot(HaveOccurred())
					})

					It("Has 1 row", func() {
						Ω(ds.NumRows()).To(Equal(1))
					})

					Context("When getting a row", func() {
						Context("When the index is correct", func() {
							It("Consistently returns the correct row", func() {
								newRow, err = ds.Row(0)
								Ω(err).ShouldNot(HaveOccurred())

								newRowAgain, err := ds.Row(0)
								Ω(err).ShouldNot(HaveOccurred())

								Ω(newRow.Features().Equals(newRowAgain.Features())).Should(BeTrue())
								Ω(newRow.Target().Equals(newRowAgain.Target())).Should(BeTrue())
							})
						})

						Context("When the index is correct", func() {
							It("Returns an error", func() {
								_, err = ds.Row(-1)
								Ω(err).Should(HaveOccurred())

								newRow, err = ds.Row(1)
								Ω(err).Should(HaveOccurred())
							})
						})
					})
				})
			})
		})

		Context("When the dataset is not empty", func() {
			BeforeEach(func() {
				err := ds.AddRowFromStrings([]string{"0.0", "hi", "mom", "94", "word"})
				Ω(err).ShouldNot(HaveOccurred())
				err = ds.AddRowFromStrings([]string{"0.0", "bye", "mom", "62", "word"})
				Ω(err).ShouldNot(HaveOccurred())
				err = ds.AddRowFromStrings([]string{"3.14", "hi", "dad", "94", "foo"})
				Ω(err).ShouldNot(HaveOccurred())
				err = ds.AddRowFromStrings([]string{"3.14", "bye", "dad", "62", "foo"})
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("Has the correct number of rows", func() {
				Ω(ds.NumRows()).To(Equal(4))
			})

			Context("When getting a row", func() {
				Context("When the index is invalid", func() {
					It("Returns an error", func() {
						_, err := ds.Row(-1)
						Ω(err).Should(HaveOccurred())

						_, err = ds.Row(4)
						Ω(err).Should(HaveOccurred())
					})
				})

				Context("When the index is valid", func() {
					It("Consistently returns the correct row", func() {
						row0, err := ds.Row(0)
						Ω(err).ShouldNot(HaveOccurred())
						row1, err := ds.Row(1)
						Ω(err).ShouldNot(HaveOccurred())
						row2, err := ds.Row(2)
						Ω(err).ShouldNot(HaveOccurred())
						row3, err := ds.Row(3)
						Ω(err).ShouldNot(HaveOccurred())

						row0Again, err := ds.Row(0)
						Ω(err).ShouldNot(HaveOccurred())
						row1Again, err := ds.Row(1)
						Ω(err).ShouldNot(HaveOccurred())
						row2Again, err := ds.Row(2)
						Ω(err).ShouldNot(HaveOccurred())
						row3Again, err := ds.Row(3)
						Ω(err).ShouldNot(HaveOccurred())

						Ω(row0.Features().Equals(row0Again.Features())).Should(BeTrue())
						Ω(row0.Target().Equals(row0Again.Target())).Should(BeTrue())
						Ω(row1.Features().Equals(row1Again.Features())).Should(BeTrue())
						Ω(row1.Target().Equals(row1Again.Target())).Should(BeTrue())
						Ω(row2.Features().Equals(row2Again.Features())).Should(BeTrue())
						Ω(row2.Target().Equals(row2Again.Target())).Should(BeTrue())
						Ω(row3.Features().Equals(row3Again.Features())).Should(BeTrue())
						Ω(row3.Target().Equals(row3Again.Target())).Should(BeTrue())

						Ω(row0.Features().Equals(row2.Features())).Should(BeTrue())
						Ω(row1.Features().Equals(row3.Features())).Should(BeTrue())
						Ω(row0.Features().Equals(row1.Features())).Should(BeFalse())

						Ω(row0.Target().Equals(row1.Target())).Should(BeTrue())
						Ω(row2.Target().Equals(row3.Target())).Should(BeTrue())
						Ω(row0.Target().Equals(row2.Target())).Should(BeFalse())
					})
				})
			})

			Context("When adding a row", func() {
				var err error

				Context("When the row's length is incorrect", func() {
					BeforeEach(func() {
						err = ds.AddRowFromStrings([]string{"0.0", "hi", "mom", "94"})
					})

					It("Returns an error", func() {
						Ω(err).Should(HaveOccurred())
					})

					It("Has the same number of rows", func() {
						Ω(ds.NumRows()).To(Equal(4))
					})
				})

				Context("When the row values don't match the column types", func() {
					BeforeEach(func() {
						err = ds.AddRowFromStrings([]string{"0.0", "hi", "mom", "shouldBeFloat", "word"})
					})

					It("Returns an error", func() {
						Ω(err).Should(HaveOccurred())
					})

					It("Has the same number of rows", func() {
						Ω(ds.NumRows()).To(Equal(4))
					})
				})

				Context("When the row is valid", func() {
					BeforeEach(func() {
						err = ds.AddRowFromStrings([]string{"0.0", "hi", "mom", "94", "word"})
					})

					It("Does not return an error", func() {
						Ω(err).ShouldNot(HaveOccurred())
					})

					It("Has 1 extra row", func() {
						Ω(ds.NumRows()).To(Equal(5))
					})

					It("Can retrieve the new row", func() {
						_, err := ds.Row(4)
						Ω(err).ShouldNot(HaveOccurred())
					})
				})
			})
		})
	})
})
