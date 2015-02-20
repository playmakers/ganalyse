package vendors_test

import (
  "../../vendors"
  "../../ganalyse"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("./Vendors/FirstDown", func() {
  const EXAMPLE_FILE = "firstdown/bikextremelite.html"

  var subject ganalyse.Product

  BeforeEach(func() {
    subject = *vendors.InspectFirstDown(load(EXAMPLE_FILE))
  })

  It("Extracts title", func() {
    Expect(subject.Name).To(Equal("Bike Xtreme Lite 73"))
  })

  It("Extracts vendor Size", func() {
    Expect(subject.Variants[0].Size).To(Equal("M"))
  })

  It("Extracts vendor Color", func() {
    Expect(subject.Variants[0].Color).To(Equal("schwarz"))
  })

  It("Extracts vendor Price", func() {
    Expect(subject.Variants[0].Price).To(Equal(170.0))
  })

})

