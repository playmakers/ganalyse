package vendors_test

import (
	"github.com/playmakers/ganalyse/lib/vendors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("./Vendors/DocA", func() {
	var exampleFile string
	var subject vendors.Product
	var variant vendors.Variant

	JustBeforeEach(func() {
		subject = *vendors.InspectDocA(loadExample(exampleFile))
		variant = *subject.DefaultVariant()
	})

	BeforeEach(func() {
		exampleFile = "doca/nikevaporjet3.0.html"
	})

	It("Extracts title", func() {
		Expect(subject.Name).To(Equal("Nike Vapor Jet 3.0 Black XLarge"))
	})

	It("Extracts vendor Size", func() {
		Expect(variant.Size).To(Equal("L"))
	})

	It("Extracts vendor Color", func() {
		Expect(variant.Color).To(Equal("Schwarz"))
	})

	It("Extracts vendor Price", func() {
		Expect(variant.Price).To(Equal(54.90))
	})

})
