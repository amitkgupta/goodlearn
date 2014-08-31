package columntype_test

import (
	. "github.com/amitkgupta/goodlearn/dataset/columntype"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stringcolumntype", func() {
	var stringColumnType ColumnType

	BeforeEach(func() {
		columnTypes, err := StringsToColumnTypes([]string{"text"})
		Ω(err).ShouldNot(HaveOccurred())
		stringColumnType = columnTypes[0]
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

	Describe("IsFloat", func() {
		It("Returns false", func() {
			Ω(stringColumnType.IsFloat()).Should(BeFalse())
		})
	})
})
