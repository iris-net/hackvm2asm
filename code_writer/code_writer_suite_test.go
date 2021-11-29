package codewriter_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCodeWriter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CodeWriter Suite")
}
