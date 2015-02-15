package vendors_test

import (
  "../../vendors"
  "../../ganalyse"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("./Vendors/UsSportshop", func() {
  const EXAMPLE_FILE = "ussportshop/nikesuperbad3.0.html"

  var subject ganalyse.Product

  BeforeEach(func() {
    subject = *vendors.InspectUsSportshop(load(EXAMPLE_FILE))
  })

  It("Extracts title", func() {
    Expect(subject.Name).To(Equal("Nike Superbad 3 Padded-Receiver Glove - navy"))
  })

  It("Extracts vendor Size", func() {
    Expect(subject.Variants[0].Size).To(Equal("S"))
  })

  It("Extracts vendor Color", func() {
    Expect(subject.Variants[0].Color).To(Equal("24"))
  })

  It("Extracts vendor Price", func() {
    Expect(subject.Variants[0].Price).To(Equal(69.9))
  })

})
