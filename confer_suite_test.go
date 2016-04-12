package confer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestConfer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Confer Suite")
}
