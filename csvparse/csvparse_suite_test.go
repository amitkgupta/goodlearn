package csvparse_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCsvparse(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Csvparse Suite")
}
