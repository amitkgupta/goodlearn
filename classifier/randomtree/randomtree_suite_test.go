package randomtree_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRandomtree(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Randomtree Suite")
}
