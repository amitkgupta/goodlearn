package gradientdescent_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGradientdescent(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gradientdescent Suite")
}
