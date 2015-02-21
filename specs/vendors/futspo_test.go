package vendors_test

import (
  "../../vendors"
  "../../ganalyse"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("./Vendors/Futspo", func() {
  var exampleFile string
  var subject ganalyse.Product
  var variant ganalyse.Variant

  JustBeforeEach(func() {
    subject = *vendors.InspectFutspo(load(exampleFile))
    variant = *subject.DefaultVariant()
  })

  BeforeEach(func() {
    exampleFile = "futspo/cutters60.html"
  })

  It("Extracts title", func() {
    Expect(subject.Name).To(Equal("Gamer S60 LB/ RB Gloves Cutters ( Shockskin Model )"))
  })

  It("Extracts vendor Size", func() {
    Expect(variant.Size).To(Equal("L"))
  })

  It("Extracts vendor Color", func() {
    Expect(variant.Color).To(Equal("charcoal/grey"))
  })

  It("Extracts vendor Price", func() {
    Expect(variant.Price).To(Equal(59.95))
  })

})
