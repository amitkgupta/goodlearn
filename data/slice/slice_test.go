package slice_test

import (
	"github.com/amitkgupta/goodlearn/data/columntype"
	"github.com/amitkgupta/goodlearn/data/slice"

	"math"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Slice", func() {
	Describe("SliceFromRawValues and Values", func() {
		var columnIndices []int
		var columnTypes []columntype.ColumnType
		var err error
		var s slice.Slice

		BeforeEach(func() {
			columnIndices = []int{1, 4, 3}
		})

		Context("When all the relevant columns store float data", func() {
			BeforeEach(func() {
				columnTypes, err = columntype.StringsToColumnTypes([]string{"x", "1.0", "1.0", "1.0", "1.0", "1.0"})
			})

			Describe("When told all entires are floats", func() {
				BeforeEach(func() {
					s, err = slice.SliceFromRawValues(true, columnIndices, columnTypes, []float64{1.2, 0, 0, 1, 4.9, 2.2})
				})

				It("Does not return an error", func() {
					Ω(err).ShouldNot(HaveOccurred())
				})

				It("Returns a float slice with the correct values", func() {
					floatSlice, ok := s.(slice.FloatSlice)
					Ω(ok).Should(BeTrue())

					Ω(floatSlice.Values()).Should(Equal([]float64{0.0, 4.9, 1.0}))
				})
			})

			Describe("When told some entires are not floats", func() {
				BeforeEach(func() {
					s, err = slice.SliceFromRawValues(false, columnIndices, columnTypes, []float64{1.2, 0, 0, 1, 4.9, 2.2})
				})

				It("Does not return an error", func() {
					Ω(err).ShouldNot(HaveOccurred())
				})

				It("Returns a mixed slice, all of whose values happen to be floats", func() {
					mixedSlice, ok := s.(slice.MixedSlice)
					Ω(ok).Should(BeTrue())

					Ω(mixedSlice.Values()).Should(Equal([]interface{}{0.0, 4.9, 1.0}))
				})
			})
		})

		Context("When some of the relevant columns store string data", func() {
			var col1val0raw, col1val1raw, col3val0raw, col3val1raw float64

			BeforeEach(func() {
				columnTypes, err = columntype.StringsToColumnTypes([]string{"1.0", "x", "1.0", "x", "1.0", "1.0"})
				Ω(err).ShouldNot(HaveOccurred())
				col1val0raw, err = columnTypes[1].PersistRawFromString("col1val0")
				Ω(err).ShouldNot(HaveOccurred())
				col1val1raw, err = columnTypes[1].PersistRawFromString("col1val1")
				Ω(err).ShouldNot(HaveOccurred())
				col3val0raw, err = columnTypes[3].PersistRawFromString("col3val0")
				Ω(err).ShouldNot(HaveOccurred())
				col3val1raw, err = columnTypes[3].PersistRawFromString("col3val1")
				Ω(err).ShouldNot(HaveOccurred())
			})

			Describe("When told all entires are floats", func() {
				BeforeEach(func() {
					rawValues := []float64{1.2, col1val0raw, 0, col3val1raw, 4.9, 2.2}
					s, err = slice.SliceFromRawValues(true, columnIndices, columnTypes, rawValues)
				})

				It("Returns an error", func() {
					Ω(err).Should(HaveOccurred())
				})
			})

			Describe("When told some entires are not floats", func() {
				Context("When the raw values corresponding to string entries are valid for their columns", func() {
					BeforeEach(func() {
						rawValues := []float64{1.2, col1val0raw, 0, col3val1raw, 4.9, 2.2}
						s, err = slice.SliceFromRawValues(false, columnIndices, columnTypes, rawValues)
					})

					It("Does not return an error", func() {
						Ω(err).ShouldNot(HaveOccurred())
					})

					It("Returns a mixed slice with the correct values", func() {
						mixedSlice, ok := s.(slice.MixedSlice)
						Ω(ok).Should(BeTrue())

						Ω(mixedSlice.Values()).Should(Equal([]interface{}{"col1val0", 4.9, "col3val1"}))
					})
				})

				Context("When some raw values corresponding to string entries are invalid for their columns", func() {
					BeforeEach(func() {
						invalidColumn1RawValue := math.Abs(col1val0raw) + math.Abs(col1val1raw) + 1
						rawValues := []float64{1.2, invalidColumn1RawValue, 0, col3val1raw, 4.9, 2.2}
						s, err = slice.SliceFromRawValues(false, columnIndices, columnTypes, rawValues)
					})

					It("Returns an error", func() {
						Ω(err).Should(HaveOccurred())
					})
				})
			})
		})
	})

	Describe("SliceFromRawValues and Equals", func() {
		var columnTypes []columntype.ColumnType
		var err error
		var s1, s2 slice.Slice

		Context("Given slices of different lengths", func() {
			BeforeEach(func() {
				columnTypes, err = columntype.StringsToColumnTypes([]string{"1.0", "1.0"})
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("returns false", func() {
				s1, _ = slice.SliceFromRawValues(true, []int{}, columnTypes, []float64{0.0, 1.0})
				s2, _ = slice.SliceFromRawValues(true, []int{0}, columnTypes, []float64{0.0, 1.0})

				Ω(s1.Equals(s2)).Should(BeFalse())
			})
		})

		Context("Given slices with different values", func() {
			BeforeEach(func() {
				columnTypes, err = columntype.StringsToColumnTypes([]string{"1.0", "1.0"})
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("returns false", func() {
				s1, _ = slice.SliceFromRawValues(true, []int{1}, columnTypes, []float64{0.0, 1.0})
				s2, _ = slice.SliceFromRawValues(true, []int{0}, columnTypes, []float64{0.0, 1.0})

				Ω(s1.Equals(s2)).Should(BeFalse())
			})
		})

		Context("Given slices with the same values in different orders", func() {
			BeforeEach(func() {
				columnTypes, err = columntype.StringsToColumnTypes([]string{"1.0", "1.0"})
				Ω(err).ShouldNot(HaveOccurred())
			})

			It("returns false", func() {
				s1, _ = slice.SliceFromRawValues(true, []int{0, 1}, columnTypes, []float64{0.0, 1.0})
				s2, _ = slice.SliceFromRawValues(true, []int{1, 0}, columnTypes, []float64{0.0, 1.0})

				Ω(s1.Equals(s2)).Should(BeFalse())
			})
		})

		Context("Given slices with the same values in the same order", func() {
			var columnTypes1, columnTypes2 []columntype.ColumnType
			var rawValues1, rawValues2 []float64

			Context("When one of the slices is a float slice, and the other is technically mixed but happens to have floats", func() {
				BeforeEach(func() {
					var helloRaw float64

					columnTypes1, err = columntype.StringsToColumnTypes([]string{"1.0", "1.0"})
					Ω(err).ShouldNot(HaveOccurred())

					columnTypes2, err = columntype.StringsToColumnTypes([]string{"x", "1.0"})
					Ω(err).ShouldNot(HaveOccurred())
					helloRaw, err = columnTypes2[0].PersistRawFromString("hello")
					Ω(err).ShouldNot(HaveOccurred())

					rawValues1 = []float64{0.0, 1.0}
					rawValues2 = []float64{helloRaw, 1.0}
				})

				It("returns true", func() {
					s1, _ = slice.SliceFromRawValues(true, []int{1}, columnTypes1, rawValues1)
					s2, _ = slice.SliceFromRawValues(false, []int{1}, columnTypes2, rawValues2)

					Ω(s1.Equals(s2)).Should(BeTrue())
				})
			})

			Context("When both slices are truly mixed slices", func() {
				BeforeEach(func() {
					var hello1Raw, hello2Raw float64

					columnTypes1, err = columntype.StringsToColumnTypes([]string{"1.0", "x"})
					Ω(err).ShouldNot(HaveOccurred())
					hello1Raw, err = columnTypes1[1].PersistRawFromString("hello")
					Ω(err).ShouldNot(HaveOccurred())
					_, err = columnTypes1[1].PersistRawFromString("goodbye")
					Ω(err).ShouldNot(HaveOccurred())

					columnTypes2, err = columntype.StringsToColumnTypes([]string{"x", "1.0"})
					Ω(err).ShouldNot(HaveOccurred())
					hello2Raw, err = columnTypes2[0].PersistRawFromString("hello")
					Ω(err).ShouldNot(HaveOccurred())

					rawValues1 = []float64{6.4, hello1Raw}
					rawValues2 = []float64{hello2Raw, 6.4}
				})

				It("returns true", func() {
					s1, _ = slice.SliceFromRawValues(false, []int{0, 1}, columnTypes1, rawValues1)
					s2, _ = slice.SliceFromRawValues(false, []int{1, 0}, columnTypes2, rawValues2)

					Ω(s1.Equals(s2)).Should(BeTrue())
				})
			})

			Context("When both slices are truly float slices", func() {
				BeforeEach(func() {
					columnTypes1, err = columntype.StringsToColumnTypes([]string{"1.0", "1.0"})
					Ω(err).ShouldNot(HaveOccurred())

					columnTypes2, err = columntype.StringsToColumnTypes([]string{"1.0", "1.0", "1.0"})
					Ω(err).ShouldNot(HaveOccurred())

					rawValues1 = []float64{0.0, 1.0}
					rawValues2 = []float64{2.0, 1.0, 0.0}
				})

				It("returns true", func() {
					s1, _ = slice.SliceFromRawValues(true, []int{0, 1}, columnTypes1, rawValues1)
					s2, _ = slice.SliceFromRawValues(true, []int{2, 1}, columnTypes2, rawValues2)

					Ω(s1.Equals(s2)).Should(BeTrue())
				})
			})
		})
	})
})
