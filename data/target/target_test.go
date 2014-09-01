package target_test

import (
	"github.com/amitkgupta/goodlearn/data/target"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Target", func() {
	Describe("Equals", func() {
		Context("Given two targets with different lengths", func() {
			It("Returns false", func() {
				Ω(target.Target{"hi", 9.0}.Equals(target.Target{"hi", 9.0, 1.2})).Should(BeFalse())
			})
		})

		Context("Given two targets with the same lengths", func() {
			Context("When they have different elements", func() {
				It("Returns false", func() {
					Ω(target.Target{"hi", 9.0}.Equals(target.Target{"mom", 8.0})).Should(BeFalse())
				})

			})

			Context("When they have the same elements in different order", func() {
				It("Returns false", func() {
					Ω(target.Target{"hi", 9.0}.Equals(target.Target{9.0, "hi"})).Should(BeFalse())
				})

			})

			Context("When they have the same elements in the same order", func() {
				It("Returns false", func() {
					Ω(target.Target{"hi", 9.0}.Equals(target.Target{"hi", 9.0})).Should(BeTrue())
				})

			})

		})
	})
})
