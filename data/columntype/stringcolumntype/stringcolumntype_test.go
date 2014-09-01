package stringcolumntype_test

import (
	"github.com/amitkgupta/goodlearn/data/columntype"
	"github.com/amitkgupta/goodlearn/data/columntype/stringcolumntype"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stringcolumntype", func() {
	var stringColumnType columntype.StringColumnType

	BeforeEach(func() {
		stringColumnType = stringcolumntype.NewStringType()
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
