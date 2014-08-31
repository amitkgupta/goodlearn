package target_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTarget(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Target Suite")
}
