package vendors_test

import (
  "../../lib/ganalyse"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
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
