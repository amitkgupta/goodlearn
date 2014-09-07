package columntype_test

import (
	"github.com/amitkgupta/goodlearn/data/columntype"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Column Type", func() {
	Describe("StringsToColumnTypes", func() {
		It("Determines whether entires are floats or strings", func() {
			columnTypes, err := columntype.StringsToColumnTypes([]string{"hi", "9.0"})
			Ω(err).ShouldNot(HaveOccurred())

			Ω(isStringColumnType(columnTypes[0])).Should(BeTrue())
			Ω(isFloatColumnType(columnTypes[1])).Should(BeTrue())
		})

		It("Handles scientific notation correctly as floats", func() {
			columnTypes, err := columntype.StringsToColumnTypes([]string{"1.0e308"})
			Ω(err).ShouldNot(HaveOccurred())

			Ω(isFloatColumnType(columnTypes[0])).Should(BeTrue())
		})

		It("Handles quoted numerals correctly as strings", func() {
			columnTypes, err := columntype.StringsToColumnTypes([]string{`"9.0"`})
			Ω(err).ShouldNot(HaveOccurred())

			Ω(isStringColumnType(columnTypes[0])).Should(BeTrue())
		})

		Context("When a string is a syntactically valid float but too big", func() {
			It("Returns an error", func() {
				_, err := columntype.StringsToColumnTypes([]string{"1.0e309"})
				Ω(err).Should(HaveOccurred())
			})
		})
	})

	Describe("Float Column Type", func() {
		var floatColumnType columntype.FloatColumnType

		BeforeEach(func() {
			columnTypes, err := columntype.StringsToColumnTypes([]string{"0.0"})
			Ω(err).ShouldNot(HaveOccurred())
			columnType := columnTypes[0]

			var ok bool
			floatColumnType, ok = columnType.(columntype.FloatColumnType)
			Ω(ok).Should(BeTrue())
		})

		Describe("ValueFromRaw", func() {
			It("Returns the given value and doesn't return an error", func() {
				value := floatColumnType.ValueFromRaw(3.14)
				Ω(value).Should(Equal(3.14))
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
	})

	Describe("Float Column Type", func() {
		var stringColumnType columntype.StringColumnType

		BeforeEach(func() {
			columnTypes, err := columntype.StringsToColumnTypes([]string{"x"})
			Ω(err).ShouldNot(HaveOccurred())
			columnType := columnTypes[0]

			var ok bool
			stringColumnType, ok = columnType.(columntype.StringColumnType)
			Ω(ok).Should(BeTrue())
		})

		Describe("ValueFromRaw and PersistRawFromString", func() {
			It("Maintains a stateful two-way map between raw values and strings", func() {
				rawHello, err := stringColumnType.PersistRawFromString("hello")
				Ω(err).ShouldNot(HaveOccurred())

				value, err := stringColumnType.ValueFromRaw(rawHello)
				Ω(value).Should(Equal("hello"))
				Ω(err).ShouldNot(HaveOccurred())

				value, err = stringColumnType.ValueFromRaw(rawHello)
				Ω(value).Should(Equal("hello"))
				Ω(err).ShouldNot(HaveOccurred())

				_, err = stringColumnType.ValueFromRaw(rawHello + 1)
				Ω(err).Should(HaveOccurred())

				rawGoodbye, err := stringColumnType.PersistRawFromString("goodbye")
				Ω(err).ShouldNot(HaveOccurred())

				value, err = stringColumnType.ValueFromRaw(rawGoodbye)
				Ω(value).Should(Equal("goodbye"))
				Ω(err).ShouldNot(HaveOccurred())

				value, err = stringColumnType.ValueFromRaw(rawHello)
				Ω(value).Should(Equal("hello"))
				Ω(err).ShouldNot(HaveOccurred())

				rawHelloNew, err := stringColumnType.PersistRawFromString("hello")
				Ω(rawHelloNew).Should(Equal(rawHello))
				Ω(err).ShouldNot(HaveOccurred())
			})
		})
	})
})

func isFloatColumnType(columnType columntype.ColumnType) bool {
	_, ok := columnType.(columntype.FloatColumnType)
	return ok
}

func isStringColumnType(columnType columntype.ColumnType) bool {
	_, ok := columnType.(columntype.StringColumnType)
	return ok
}
