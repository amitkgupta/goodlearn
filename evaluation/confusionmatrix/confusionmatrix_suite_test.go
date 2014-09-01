package confusionmatrix_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestConfusionmatrix(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Confusionmatrix Suite")
}
