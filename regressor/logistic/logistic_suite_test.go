package logistic_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLogistic(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Logistic Suite")
}
