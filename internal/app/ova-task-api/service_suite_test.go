package api_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOvaTaskApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OvaTaskApi Suite")
}
