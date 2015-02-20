package vendors_test

import (
  "../../vendors"
  "../../ganalyse"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("./Vendors/Potsdam", func() {
  Context("when no color given", func() {
    const EXAMPLE_FILE = "potsdam/riddellrevospeedm-l.html"

    var subject ganalyse.Product

    BeforeEach(func() {
      subject = *vendors.InspectPotsdam(load(EXAMPLE_FILE))
    })

    It("Extracts title", func() {
      Expect(subject.Name).To(Equal("*Revolution Speed Helm (ohne Facemasks)"))
    })

    It("Extracts vendor Size", func() {
      Expect(subject.Variants[0].Size).To(Equal("XL"))
    })

    It("Extracts vendor Color", func() {
      Expect(subject.Variants[0].Color).To(Equal("51"))
    })

    It("Extracts vendor Price", func() {
      Expect(subject.Variants[0].Price).To(Equal(289.95))
    })
  })

  Context("when no color given", func() {
    const EXAMPLE_FILE = "potsdam/nikesuperbad3.0.html"

    var subject ganalyse.Product

    BeforeEach(func() {
      subject = *vendors.InspectPotsdam(load(EXAMPLE_FILE))
    })

    It("Extracts title", func() {
      Expect(subject.Name).To(Equal("Nike Super Bad 3.0 schwarz"))
    })

    It("Extracts vendor Size", func() {
      Expect(subject.Variants[0].Size).To(Equal("M"))
    })

    It("Extracts vendor Color", func() {
      Expect(subject.Variants[0].Color).To(Equal("schwarz"))
    })

    It("Extracts vendor Price", func() {
      Expect(subject.Variants[0].Price).To(Equal(69.95))
    })
  })

})
