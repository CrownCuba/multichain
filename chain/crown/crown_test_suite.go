package crown_test

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	"testing"
)

func TestCrown(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Crown Suite")
}