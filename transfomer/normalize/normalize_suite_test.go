package normalize_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestNormalize(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Normalize Suite")
}
