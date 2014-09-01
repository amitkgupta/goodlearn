package stochasticgradientdescent_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestStochasticgradientdescent(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stochasticgradientdescent Suite")
}
