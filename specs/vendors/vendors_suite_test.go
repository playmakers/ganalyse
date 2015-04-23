package vendors_test

import (
	"io/ioutil"
	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func loadExample(path string) []byte {
	path = "../../examples/" + path
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Couldn't load URL")
	}
	return data
}

func TestVendors(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Vendors Suite")
}
