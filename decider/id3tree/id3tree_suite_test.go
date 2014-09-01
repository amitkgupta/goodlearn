package id3tree_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestId3tree(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Id3tree Suite")
}
