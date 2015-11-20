package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoBad(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoBad Suite")
}
