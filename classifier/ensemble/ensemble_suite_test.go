package ensemble_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEnsemble(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ensemble Suite")
}
