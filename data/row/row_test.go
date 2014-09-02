package row_test

import (
	"github.com/amitkgupta/goodlearn/data/row"
	"github.com/amitkgupta/goodlearn/data/target"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Row", func() {
	Describe("NewRow", func() {
		var allFeaturesFloats bool

		Context("When told that all features are floats", func() {
			BeforeEach(func() {
				allFeaturesFloats = true
			})

			It("Returns a FloatFeatureRow", func() {
				row := row.NewRow(allFeaturesFloats, target.Target{}, []float64{0})
				Ω(isFloatFeatureRow(row)).Should(BeTrue())
			})
		})

		Context("When told that not all features are floats", func() {
			BeforeEach(func() {
				allFeaturesFloats = false
			})

			It("Returns a MixedFeatureRow", func() {
				row := row.NewRow(allFeaturesFloats, target.Target{}, []float64{0})
				Ω(isMixedFeatureRow(row)).Should(BeTrue())
			})
		})
	})
})

func isFloatFeatureRow(r row.Row) bool {
	_, ok := r.(row.FloatFeatureRow)
	return ok
}

func isMixedFeatureRow(r row.Row) bool {
	_, ok := r.(row.MixedFeatureRow)
	return ok
}
