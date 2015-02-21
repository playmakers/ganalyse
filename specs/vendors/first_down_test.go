package vendors_test

import (
  "../../vendors"
  "../../ganalyse"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("./Vendors/FirstDown", func() {
  var exampleFile string
  var subject ganalyse.Product
  var variant ganalyse.Variant

  JustBeforeEach(func() {
    subject = *vendors.InspectFirstDown(load(exampleFile))
    variant = *subject.DefaultVariant()
  })

  BeforeEach(func() {
    exampleFile = "firstdown/bikextremelite.html"
  })

  It("Extracts title", func() {
    Expect(subject.Name).To(Equal("Bike Xtreme Lite 73"))
  })

  It("Extracts vendor Size", func() {
    Expect(variant.Size).To(Equal("L"))
  })

  It("Extracts vendor Color", func() {
    Expect(variant.Color).To(Equal("Schwarz"))
  })

  It("Extracts vendor Price", func() {
    Expect(variant.Price).To(Equal(170.0))
  })

})

