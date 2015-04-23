package vendors_test

import (
	"github.com/playmakers/ganalyse/lib/vendors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("./Vendors/Meyer", func() {
	var exampleFile string
	var subject vendors.Product
	var variant vendors.Variant

	JustBeforeEach(func() {
		subject = *vendors.InspectMeyer(loadExample(exampleFile))
		variant = *subject.DefaultVariant()
	})

	BeforeEach(func() {
		exampleFile = "meyer/xenithepic.html"
	})

	It("Extracts title", func() {
		Expect(subject.Name).To(Equal("EPIC   XENITH Epic Football Helmet"))
	})

	It("Extracts vendor Size", func() {
		Expect(variant.Size).To(Equal("L"))
	})

	It("Extracts vendor Color", func() {
		Expect(variant.Color).To(Equal("Schwarz"))
	})

	It("Extracts vendor Price", func() {
		Expect(variant.Price).To(Equal(360.0))
	})

})
