package dataset_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDataset(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dataset Suite")
}
