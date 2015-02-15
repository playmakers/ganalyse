package vendors_test

import (
	"../../vendors"
  "../../ganalyse"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("./Vendors/1a", func() {
  const EXAMPLE_FILE  = "1a/bikextremelite.html"

  Context("when no Variants are available", func() {
    var subject ganalyse.Product

    BeforeEach(func() {
      subject = *vendors.Inspect1A(load(EXAMPLE_FILE))
    })

    It("Extracts title", func() {
      Expect(subject.Name).To(Equal("X-treme Lite RB/DB"))
    })
  })

  Context("when Variants are available", func() {
    const EXAMPLE_FILE2 = "1a/cutters60.html"

    var subject ganalyse.Variant

    BeforeEach(func() {
      subject = vendors.Inspect1A(load(EXAMPLE_FILE2)).Variants[0]
    })

    It("Extracts vendor Size", func() {
      Expect(subject.Size).To(Equal("S"))
    })

    It("Extracts vendor Color", func() {
      Expect(subject.Color).To(Equal("Dunkel Blau"))
    })

    It("Extracts vendor Price", func() {
      Expect(subject.Price).To(Equal(57.5))
    })
  })

})
