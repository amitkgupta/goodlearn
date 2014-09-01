package columntype_test

import (
	"github.com/amitkgupta/goodlearn/data/columntype"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Columntype", func() {
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
})

func isFloatColumnType(columnType columntype.ColumnType) bool {
	_, ok := columnType.(columntype.FloatColumnType)
	return ok
}

func isStringColumnType(columnType columntype.ColumnType) bool {
	_, ok := columnType.(columntype.StringColumnType)
	return ok
}
