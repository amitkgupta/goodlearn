package columntype_test

import (
	"github.com/amitkgupta/goodlearn/data/columntype"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Columntype", func() {
	Describe("StringsToColumnTypes", func() {
		It("Determines whether entires are floats or not", func() {
			columnTypes, err := columntype.StringsToColumnTypes([]string{"hi", "9.0"})
			Ω(err).ShouldNot(HaveOccurred())

			Ω(columnTypes[0].IsFloat()).Should(BeFalse())
			Ω(columnTypes[1].IsFloat()).Should(BeTrue())
		})

		It("Handles scientific notation correctly as floats", func() {
			columnTypes, err := columntype.StringsToColumnTypes([]string{"1.0e308"})
			Ω(err).ShouldNot(HaveOccurred())

			Ω(columnTypes[0].IsFloat()).Should(BeTrue())
		})

		It("Handles quoted numerals correctly as non-floats", func() {
			columnTypes, err := columntype.StringsToColumnTypes([]string{`"9.0"`})
			Ω(err).ShouldNot(HaveOccurred())

			Ω(columnTypes[0].IsFloat()).Should(BeFalse())
		})

		Context("When a string is a syntactically valid float but too big", func() {
			It("Returns an error", func() {
				_, err := columntype.StringsToColumnTypes([]string{"1.0e309"})
				Ω(err).Should(HaveOccurred())
			})
		})
	})
})
