package vectorutilities_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestVectorutilities(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Vectorutilities Suite")
}
