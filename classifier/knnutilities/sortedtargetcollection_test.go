package knnutilities_test

import (
	. "github.com/amitkgupta/goodlearn/classifier/knnutilities"
	"github.com/amitkgupta/goodlearn/dataset/target"

	"math"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SortedTargetCollection", func() {
	Describe("NewKNNTargetCollection", func() {
		Context("W given a non-positive k value", func() {
			It("Does not return an error", func() {
				_, err := NewKNNTargetCollection(0)
				Ω(err).Should(HaveOccurred())

				_, err = NewKNNTargetCollection(-1)
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("When given a positive k value", func() {
			It("Returns an error", func() {
				_, err := NewKNNTargetCollection(1)
				Ω(err).ShouldNot(HaveOccurred())

				_, err = NewKNNTargetCollection(50)
				Ω(err).ShouldNot(HaveOccurred())
			})
		})
	})

	Describe("Insert and MaxDistance", func() {
		var stc SortedTargetCollection

		BeforeEach(func() {
			stc, _ = NewKNNTargetCollection(2)
		})

		Context("Before the collection is full", func() {
			It("Does not return errors", func() {
				err := stc.Insert(target.Target{}, 1.0)
				Ω(stc.MaxDistance()).Should(Equal(math.MaxFloat64))
				Ω(err).ShouldNot(HaveOccurred())

				err = stc.Insert(target.Target{}, 3.0)
				Ω(stc.MaxDistance()).Should(Equal(3.0))
				Ω(err).ShouldNot(HaveOccurred())
			})
		})

		Context("When the collection is full", func() {
			initialMax := 3.0
			initialMin := initialMax - 3

			BeforeEach(func() {
				err := stc.Insert(target.Target{}, initialMin)
				Ω(err).ShouldNot(HaveOccurred())

				err = stc.Insert(target.Target{}, initialMax)
				Ω(err).ShouldNot(HaveOccurred())
			})

			Context("When inserting an element with too large a distance", func() {
				It("Returns an error", func() {
					err := stc.Insert(target.Target{}, initialMax)
					Ω(stc.MaxDistance()).Should(Equal(initialMax))
					Ω(err).Should(HaveOccurred())

					err = stc.Insert(target.Target{}, initialMax+2)
					Ω(stc.MaxDistance()).Should(Equal(initialMax))
					Ω(err).Should(HaveOccurred())
				})
			})

			Context("When inserting an element with a small enough distance", func() {
				It("Does not return an error", func() {
					err := stc.Insert(target.Target{}, initialMax-2)
					Ω(stc.MaxDistance()).Should(Equal(initialMax - 2))
					Ω(err).ShouldNot(HaveOccurred())
				})

				It("Behaves as though the distance threshold has properly decreased", func() {
					err := stc.Insert(target.Target{}, initialMin)
					Ω(stc.MaxDistance()).Should(Equal(initialMin))
					Ω(err).ShouldNot(HaveOccurred())

					err = stc.Insert(target.Target{}, initialMin)
					Ω(stc.MaxDistance()).Should(Equal(initialMin))
					Ω(err).Should(HaveOccurred())

					err = stc.Insert(target.Target{}, initialMin-2)
					Ω(stc.MaxDistance()).Should(Equal(initialMin))
					Ω(err).ShouldNot(HaveOccurred())

					err = stc.Insert(target.Target{}, initialMin-1)
					Ω(stc.MaxDistance()).Should(Equal(initialMin - 1))
					Ω(err).ShouldNot(HaveOccurred())

					err = stc.Insert(target.Target{}, initialMin-1)
					Ω(stc.MaxDistance()).Should(Equal(initialMin - 1))
					Ω(err).Should(HaveOccurred())
				})
			})
		})
	})

	Describe("Vote", func() {
		var stc SortedTargetCollection

		Context("When the collection is empty", func() {
			BeforeEach(func() {
				stc, _ = NewKNNTargetCollection(1)
			})

			It("Returns an error", func() {
				_, err := stc.Vote()
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("When the collection is not empty", func() {
			Context("When the collection has fewer than k candidates", func() {
				BeforeEach(func() {
					stc, _ = NewKNNTargetCollection(2)
					err := stc.Insert(target.Target{"hi"}, 23.0)
					Ω(err).ShouldNot(HaveOccurred())
				})

				It("Does not return an error", func() {
					_, err := stc.Vote()
					Ω(err).ShouldNot(HaveOccurred())
				})

				It("Chooses the winner", func() {
					winner, _ := stc.Vote()
					Ω(winner.Equals(target.Target{"hi"})).Should(BeTrue())
				})
			})

			Context("When there is a clear winner", func() {
				BeforeEach(func() {
					stc, _ = NewKNNTargetCollection(3)
					err := stc.Insert(target.Target{"hi"}, 23.0)
					Ω(err).ShouldNot(HaveOccurred())
					err = stc.Insert(target.Target{"mom"}, 40.0)
					Ω(err).ShouldNot(HaveOccurred())
					err = stc.Insert(target.Target{"mom"}, 42.0)
					Ω(err).ShouldNot(HaveOccurred())
				})

				It("Chooses the winner", func() {
					winner, _ := stc.Vote()
					Ω(winner.Equals(target.Target{"mom"})).Should(BeTrue())
				})
			})

			Context("When there is a tie", func() {
				BeforeEach(func() {
					stc, _ = NewKNNTargetCollection(5)
					err := stc.Insert(target.Target{"hi", "hello"}, 23.0)
					Ω(err).ShouldNot(HaveOccurred())
					err = stc.Insert(target.Target{"mom", "dad"}, 1.0)
					Ω(err).ShouldNot(HaveOccurred())
					err = stc.Insert(target.Target{"mom", "dad"}, 42.0)
					Ω(err).ShouldNot(HaveOccurred())
					err = stc.Insert(target.Target{"hi", "hello"}, 24.0)
					Ω(err).ShouldNot(HaveOccurred())
					err = stc.Insert(target.Target{"mom", "loser"}, 0.5)
					Ω(err).ShouldNot(HaveOccurred())
				})

				It("Chooses from amongst the tied candidates", func() {
					winner, _ := stc.Vote()
					candidate1 := target.Target{"mom", "dad"}
					candidate2 := target.Target{"hi", "hello"}

					Ω(winner.Equals(candidate1) || winner.Equals(candidate2)).Should(BeTrue())
				})
			})
		})
	})
})
