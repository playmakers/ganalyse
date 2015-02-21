package vendors_test

import (
  "../../vendors"
  "../../ganalyse"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("./Vendors/Forelle", func() {
  var exampleFile string
  var subject ganalyse.Product
  var variant ganalyse.Variant

  JustBeforeEach(func() {
    subject = *vendors.InspectForelle(load(exampleFile))
    variant = *subject.DefaultVariant()
  })

  BeforeEach(func() {
    exampleFile = "forelle/nikesuperbad3.0.html"
  })

  It("Extracts title", func() {
    Expect(subject.Name).To(Equal("Nike Super Bad 3.0"))
  })

  It("Extracts vendor Size", func() {
    Expect(variant.Size).To(Equal("L"))
  })

  It("Extracts vendor Color", func() {
    Expect(variant.Color).To(Equal("1"))
  })

  It("Extracts vendor Price", func() {
    Expect(variant.Price).To(Equal(79.0))
  })

})
