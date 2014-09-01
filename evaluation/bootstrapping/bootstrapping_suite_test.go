package bootstrapping_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBootstrapping(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bootstrapping Suite")
}
