package vendors_test

import (
  "../../vendors"
  "../../ganalyse"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("./Vendors/Futspo", func() {
  const EXAMPLE_FILE = "futspo/cutters60.html"

  var subject ganalyse.Product

  BeforeEach(func() {
    subject = *vendors.InspectFutspo(load(EXAMPLE_FILE))
  })

  It("Extracts title", func() {
    Expect(subject.Name).To(Equal("Gamer S60 LB/ RB Gloves Cutters ( Shockskin Model )"))
  })

  It("Extracts vendor Size", func() {
    Expect(subject.Variants[0].Size).To(Equal("XL"))
  })

  It("Extracts vendor Color", func() {
    Expect(subject.Variants[0].Color).To(Equal("schwarz"))
  })

  It("Extracts vendor Price", func() {
    Expect(subject.Variants[0].Price).To(Equal(59.95))
  })

})
