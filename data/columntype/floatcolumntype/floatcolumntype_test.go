package floatcolumntype_test

import (
	"github.com/amitkgupta/goodlearn/data/columntype"
	"github.com/amitkgupta/goodlearn/data/columntype/floatcolumntype"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Floatcolumntype", func() {
	var floatColumnType columntype.FloatColumnType

	BeforeEach(func() {
		floatColumnType = floatcolumntype.NewFloatType()
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
})
