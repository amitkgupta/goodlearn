package knnutilities_test

import (
	"github.com/amitkgupta/goodlearn/classifier/knn/knnutilities"
	"github.com/amitkgupta/goodlearn/data/columntype"
	"github.com/amitkgupta/goodlearn/data/slice"

	"math"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SortedTargetCollection", func() {
	Describe("Insert and MaxDistance", func() {
		var stc knnutilities.SortedTargetCollection
		var target slice.Slice
		var err error

		BeforeEach(func() {
			stc = knnutilities.NewKNNTargetCollection(2)
			target, err = slice.SliceFromRawValues(true, []int{}, []columntype.ColumnType{}, []float64{})
			Ω(err).ShouldNot(HaveOccurred())
		})

		Context("Before the collection is full", func() {
			It("The MaxDistance should be +Inf", func() {
				Ω(stc.MaxDistance()).Should(Equal(math.MaxFloat64))

				stc.Insert(target, 1.0)
				Ω(stc.MaxDistance()).Should(Equal(math.MaxFloat64))
			})
		})

		Context("When the collection is full", func() {
			initialMax := 3.0
			initialMin := initialMax - 3

			BeforeEach(func() {
				stc.Insert(target, initialMax)
				stc.Insert(target, initialMin)
			})

			It("The MaxDistance should be the distance of the 'farthest' target", func() {
				Ω(stc.MaxDistance()).Should(Equal(initialMax))
			})

			Context("When inserting an element with whose distance is larger than MaxDistance", func() {
				It("The MaxDistance should not change", func() {
					Ω(stc.MaxDistance()).Should(Equal(initialMax))

					stc.Insert(target, initialMax+2)
					Ω(stc.MaxDistance()).Should(Equal(initialMax))
				})
			})

			Context("When inserting an elements with a small enough distances", func() {
				It("Behaves as though the distance threshold has properly decreased", func() {
					Ω(stc.MaxDistance()).Should(Equal(initialMax))

					stc.Insert(target, initialMin)
					Ω(stc.MaxDistance()).Should(Equal(initialMin))

					stc.Insert(target, initialMin)
					Ω(stc.MaxDistance()).Should(Equal(initialMin))

					stc.Insert(target, initialMin-2)
					Ω(stc.MaxDistance()).Should(Equal(initialMin))

					stc.Insert(target, initialMin-1)
					Ω(stc.MaxDistance()).Should(Equal(initialMin - 1))

					stc.Insert(target, initialMin-1)
					Ω(stc.MaxDistance()).Should(Equal(initialMin - 1))
				})
			})
		})
	})

	Describe("Vote", func() {
		var stc knnutilities.SortedTargetCollection

		Context("When the collection is not empty", func() {
			var target1, target2, target3 slice.Slice

			BeforeEach(func() {
				stc = knnutilities.NewKNNTargetCollection(5)

				columnTypes, err := columntype.StringsToColumnTypes([]string{"1.0"})
				Ω(err).ShouldNot(HaveOccurred())

				raw1, err := columnTypes[0].PersistRawFromString("1.0")
				Ω(err).ShouldNot(HaveOccurred())

				raw2, err := columnTypes[0].PersistRawFromString("2.0")
				Ω(err).ShouldNot(HaveOccurred())

				raw3, err := columnTypes[0].PersistRawFromString("3.0")
				Ω(err).ShouldNot(HaveOccurred())

				target1, err = slice.SliceFromRawValues(true, []int{0}, columnTypes, []float64{raw1})
				target2, err = slice.SliceFromRawValues(true, []int{0}, columnTypes, []float64{raw2})
				target3, err = slice.SliceFromRawValues(true, []int{0}, columnTypes, []float64{raw3})
			})

			Context("When the collection has fewer than k candidates", func() {
				BeforeEach(func() {
					stc.Insert(target1, 23.0)
				})

				It("Chooses the winner", func() {
					winner := stc.Vote()
					Ω(target1.Equals(winner)).Should(BeTrue())
				})
			})

			Context("When there is a clear winner", func() {
				BeforeEach(func() {
					stc.Insert(target1, 23.0)
					stc.Insert(target2, 40.0)
					stc.Insert(target2, 50.0)
					stc.Insert(target2, 90.0)
					stc.Insert(target3, 10.0)
				})

				It("Chooses the winner", func() {
					winner := stc.Vote()
					Ω(target2.Equals(winner)).Should(BeTrue())
				})
			})

			Context("When there is a tie", func() {
				BeforeEach(func() {
					stc = knnutilities.NewKNNTargetCollection(5)
					stc.Insert(target1, 23.0)
					stc.Insert(target2, 1.0)
					stc.Insert(target2, 42.0)
					stc.Insert(target3, 24.0)
					stc.Insert(target3, 0.5)
				})

				It("Chooses from amongst the tied candidates", func() {
					winner := stc.Vote()

					Ω(winner.Equals(target2) || winner.Equals(target3)).Should(BeTrue())
				})
			})
		})
	})
})
