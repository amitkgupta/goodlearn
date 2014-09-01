package inmemoryrow_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestInmemoryrow(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Inmemoryrow Suite")
}
