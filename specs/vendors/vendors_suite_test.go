package vendors_test

import (
	"github.com/playmakers/ganalyse/lib/ganalyse"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func load(filename string) []byte {
	filename = "../../examples/" + filename
	return LoadFile(filename)
}

func TestVendors(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Vendors Suite")
}
