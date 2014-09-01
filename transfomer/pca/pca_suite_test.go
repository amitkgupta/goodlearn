package pca_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPca(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pca Suite")
}
