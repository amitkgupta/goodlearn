package knnutilities_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestKnnutilities(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Knnutilities Suite")
}
