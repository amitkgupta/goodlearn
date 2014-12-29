package matrixparse_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMatrixparse(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Matrixparse Suite")
}
