package vendors_test

import (
  "../../lib/vendors"
  "../../lib/ganalyse"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("./Vendors/BossHogg", func() {
  var exampleFile string
  var subject ganalyse.Product
  var variant ganalyse.Variant

  JustBeforeEach(func() {
    subject = *vendors.InspectBossHogg(load(exampleFile))
    variant = *subject.DefaultVariant()
  })

  BeforeEach(func() {
    exampleFile = "bosshogg/riddell360m-l.html"
  })

  It("Extracts title", func() {
    Expect(subject.Name).To(Equal("Riddell 360 Helm"))
  })

  It("Extracts vendor Size", func() {
    Expect(variant.Size).To(Equal("L"))
  })

  It("Extracts vendor Color", func() {
    Expect(variant.Color).To(Equal("Schwarz"))
  })

  It("Extracts vendor Price", func() {
    Expect(variant.Price).To(Equal(385.0))
  })

})
