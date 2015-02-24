package vendors_test

import (
  "../../lib/vendors"
  "../../lib/ganalyse"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("./Vendors/DocA", func() {
  var exampleFile string
  var subject ganalyse.Product
  var variant ganalyse.Variant

  JustBeforeEach(func() {
    subject = *vendors.InspectDocA(load(exampleFile))
    variant = *subject.DefaultVariant()
  })

  BeforeEach(func() {
    exampleFile = "doca/nikesuperbad3.0.html"
  })

  It("Extracts title", func() {
    Expect(subject.Name).To(Equal("Nike Superbad 3.0 Black"))
  })

  It("Extracts vendor Size", func() {
    Expect(variant.Size).To(Equal("L"))
  })

  It("Extracts vendor Color", func() {
    Expect(variant.Color).To(Equal("Schwarz"))
  })

  It("Extracts vendor Price", func() {
    Expect(variant.Price).To(Equal(59.90))
  })

})
