package knn_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestKnn(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Knn Suite")
}
