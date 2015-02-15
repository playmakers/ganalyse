package vendors_test

import (
	"../../vendors"
  "../../ganalyse"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("./Vendors/DocA", func() {
  const EXAMPLE_FILE = "doca/nikesuperbad3.0.html"

  var subject ganalyse.Product

  BeforeEach(func() {
    subject = *vendors.InspectDocA(load(EXAMPLE_FILE))
  })

  It("Extracts title", func() {
    Expect(subject.Name).To(Equal("Nike Superbad 3.0 Black"))
  })

  It("Extracts vendor Size", func() {
    Expect(subject.Variants[0].Size).To(Equal("L"))
  })

  It("Extracts vendor Color", func() {
    Expect(subject.Variants[0].Color).To(Equal("schwarz"))
  })

  It("Extracts vendor Price", func() {
    Expect(subject.Variants[0].Price).To(Equal(59.90))
  })

})
