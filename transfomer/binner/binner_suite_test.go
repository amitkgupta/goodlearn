package binner_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBinner(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Binner Suite")
}
