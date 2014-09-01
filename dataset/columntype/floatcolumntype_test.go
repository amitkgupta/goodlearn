package columntype_test

import (
	. "github.com/amitkgupta/goodlearn/dataset/columntype"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Floatcolumntype", func() {
	var floatColumnType ColumnType

	BeforeEach(func() {
		columnTypes, err := StringsToColumnTypes([]string{"1.0"})
		Ω(err).ShouldNot(HaveOccurred())
		floatColumnType = columnTypes[0]
	})

	Describe("ValueFromRaw", func() {
		It("Returns the given value and doesn't return an error", func() {
			value, err := floatColumnType.ValueFromRaw(3.14)
			Ω(value).Should(Equal(3.14))
			Ω(err).ShouldNot(HaveOccurred())
		})
	})

	Describe("PersistRawFromString", func() {
		Context("Given a string representing a valid float", func() {
			It("Returns the float and no error", func() {
				value, err := floatColumnType.PersistRawFromString("3.14")
				Ω(value).Should(Equal(3.14))
				Ω(err).ShouldNot(HaveOccurred())
			})
		})

		Context("Given a string representing an invalid float", func() {
			It("Returns an error", func() {
				_, err := floatColumnType.PersistRawFromString("pi")
				Ω(err).Should(HaveOccurred())

			})
		})
	})

	Describe("IsFloat", func() {
		It("Returns true", func() {
			Ω(floatColumnType.IsFloat()).Should(BeTrue())
		})
	})
})
