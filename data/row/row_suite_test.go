package row_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRow(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Row Suite")
}
