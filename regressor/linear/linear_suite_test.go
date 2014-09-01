package linear_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLinear(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Linear Suite")
}
