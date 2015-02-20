package vendors_test

import (
  "../../vendors"
  "../../ganalyse"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("./Vendors/SportsAndCheer", func() {
  const EXAMPLE_FILE = "sportsandcheer/cutters60.html"

  var subject ganalyse.Product

  BeforeEach(func() {
    subject = *vendors.InspectSportsAndCheer(load(EXAMPLE_FILE))
  })

  It("Extracts title", func() {
    Expect(subject.Name).To(Equal("Cutters The Gamer Football Handschuhe"))
  })

  It("Extracts vendor Size", func() {
    Expect(subject.Variants[0].Size).To(Equal("S"))
  })

  It("Extracts vendor Color", func() {
    Expect(subject.Variants[0].Color).To(Equal("Schwarz"))
  })

  It("Extracts vendor Price", func() {
    Expect(subject.Variants[0].Price).To(Equal(55.0))
  })

})
