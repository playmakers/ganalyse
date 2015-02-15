package vendors_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

  "../../ganalyse"
	"testing"
)

func load(filename string) []byte {
  filename = "../../examples/" + filename
  return ganalyse.LoadFile(filename)
}

func TestVendors(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Vendors Suite")
}
