package naivebayes_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestNaivebayes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Naivebayes Suite")
}
