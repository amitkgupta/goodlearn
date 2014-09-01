package kmeans_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestKmeans(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kmeans Suite")
}
