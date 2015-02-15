package vendors_test

import (
	"../../vendors"
  "../../ganalyse"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("./Vendors/BossHogg", func() {
  const EXAMPLE_FILE = "bosshogg/riddell360m-l.html"

  var subject ganalyse.Product

  BeforeEach(func() {
    subject = *vendors.InspectBossHogg(load(EXAMPLE_FILE))
  })

  It("Extracts title", func() {
    Expect(subject.Name).To(Equal("Riddell 360 Helm"))
  })

  It("Extracts vendor Size", func() {
    Expect(subject.Variants[0].Size).To(Equal("L"))
  })

  It("Extracts vendor Color", func() {
    Expect(subject.Variants[0].Color).To(Equal("schwarz"))
  })

  It("Extracts vendor Price", func() {
    Expect(subject.Variants[0].Price).To(Equal(385.0))
  })

})
