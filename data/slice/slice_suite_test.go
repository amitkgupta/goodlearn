package slice_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSlice(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Slice Suite")
}
