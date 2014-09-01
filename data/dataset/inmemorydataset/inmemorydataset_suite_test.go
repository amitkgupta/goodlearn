package inmemorydataset_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestInmemorydataset(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Inmemorydataset Suite")
}
