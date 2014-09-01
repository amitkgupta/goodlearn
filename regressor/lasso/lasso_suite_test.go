package lasso_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLasso(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lasso Suite")
}
