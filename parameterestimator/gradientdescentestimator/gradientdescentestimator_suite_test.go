package gradientdescentestimator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGradientdescentestimator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gradientdescentestimator Suite")
}
