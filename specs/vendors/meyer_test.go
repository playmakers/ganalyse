package vendors_test

import (
	"github.com/playmakers/ganalyse/lib/ganalyse"
	"github.com/playmakers/ganalyse/lib/vendors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("./Vendors/Meyer", func() {
	var exampleFile string
	var subject Product
	var variant Variant

	JustBeforeEach(func() {
		subject = *InspectMeyer(load(exampleFile))
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
		Expect(variant.Price).To(Equal(340.0))
	})

})
