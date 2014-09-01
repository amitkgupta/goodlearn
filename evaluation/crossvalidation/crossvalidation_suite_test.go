package crossvalidation_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCrossvalidation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Crossvalidation Suite")
}
