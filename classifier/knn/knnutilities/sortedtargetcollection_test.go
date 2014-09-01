package knnutilities_test

import (
	"github.com/amitkgupta/goodlearn/classifier/knn/knnutilities"
	"github.com/amitkgupta/goodlearn/data/target"

	"math"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SortedTargetCollection", func() {
	Describe("Insert and MaxDistance", func() {
		var stc knnutilities.SortedTargetCollection

		BeforeEach(func() {
			stc = knnutilities.NewKNNTargetCollection(2)
		})

		Context("Before the collection is full", func() {
			It("The MaxDistance should be +Inf", func() {
				Ω(stc.MaxDistance()).Should(Equal(math.MaxFloat64))

				stc.Insert(target.Target{}, 1.0)
				Ω(stc.MaxDistance()).Should(Equal(math.MaxFloat64))
			})
		})

		Context("When the collection is full", func() {
			initialMax := 3.0
			initialMin := initialMax - 3

			BeforeEach(func() {
				stc.Insert(target.Target{}, initialMax)
				stc.Insert(target.Target{}, initialMin)
			})

			It("The MaxDistance should be the distance of the 'farthest' target", func() {
				Ω(stc.MaxDistance()).Should(Equal(initialMax))
			})

			Context("When inserting an element with whose distance is larger than MaxDistance", func() {
				It("The MaxDistance should not change", func() {
					Ω(stc.MaxDistance()).Should(Equal(initialMax))

					stc.Insert(target.Target{}, initialMax+2)
					Ω(stc.MaxDistance()).Should(Equal(initialMax))
				})
			})

			Context("When inserting an elements with a small enough distances", func() {
				It("Behaves as though the distance threshold has properly decreased", func() {
					Ω(stc.MaxDistance()).Should(Equal(initialMax))

					stc.Insert(target.Target{}, initialMin)
					Ω(stc.MaxDistance()).Should(Equal(initialMin))

					stc.Insert(target.Target{}, initialMin)
					Ω(stc.MaxDistance()).Should(Equal(initialMin))

					stc.Insert(target.Target{}, initialMin-2)
					Ω(stc.MaxDistance()).Should(Equal(initialMin))

					stc.Insert(target.Target{}, initialMin-1)
					Ω(stc.MaxDistance()).Should(Equal(initialMin - 1))

					stc.Insert(target.Target{}, initialMin-1)
					Ω(stc.MaxDistance()).Should(Equal(initialMin - 1))
				})
			})
		})
	})

	Describe("Vote", func() {
		var stc knnutilities.SortedTargetCollection

		Context("When the collection is not empty", func() {
			Context("When the collection has fewer than k candidates", func() {
				BeforeEach(func() {
					stc = knnutilities.NewKNNTargetCollection(2)
					stc.Insert(target.Target{"hi"}, 23.0)
				})

				It("Chooses the winner", func() {
					Ω(stc.Vote().Equals(target.Target{"hi"})).Should(BeTrue())
				})
			})

			Context("When there is a clear winner", func() {
				BeforeEach(func() {
					stc = knnutilities.NewKNNTargetCollection(3)
					stc.Insert(target.Target{"hi"}, 23.0)
					stc.Insert(target.Target{"mom"}, 40.0)
					stc.Insert(target.Target{"mom"}, 42.0)
				})

				It("Chooses the winner", func() {
					Ω(stc.Vote().Equals(target.Target{"mom"})).Should(BeTrue())
				})
			})

			Context("When there is a tie", func() {
				BeforeEach(func() {
					stc = knnutilities.NewKNNTargetCollection(5)
					stc.Insert(target.Target{"hi", "hello"}, 23.0)
					stc.Insert(target.Target{"mom", "dad"}, 1.0)
					stc.Insert(target.Target{"mom", "dad"}, 42.0)
					stc.Insert(target.Target{"hi", "hello"}, 24.0)
					stc.Insert(target.Target{"mom", "loser"}, 0.5)
				})

				It("Chooses from amongst the tied candidates", func() {
					winner := stc.Vote()
					candidate1 := target.Target{"mom", "dad"}
					candidate2 := target.Target{"hi", "hello"}

					Ω(winner.Equals(candidate1) || winner.Equals(candidate2)).Should(BeTrue())
				})
			})
		})
	})
})
