package vendors_test

import (
	"github.com/playmakers/ganalyse/lib/vendors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("./Vendors/Potsdam", func() {
	var exampleFile string
	var subject vendors.Product
	var variant vendors.Variant

	JustBeforeEach(func() {
		subject = *vendors.InspectPotsdam(load(exampleFile))
		variant = *subject.DefaultVariant()
	})

	Context("when no color given", func() {
		BeforeEach(func() {
			exampleFile = "potsdam/riddellrevospeedm-l.html"
		})

		It("Extracts title", func() {
			Expect(subject.Name).To(Equal("*Revolution Speed Helm (ohne Facemasks)"))
		})

		It("Extracts vendor Size", func() {
			Expect(variant.Size).To(Equal("L"))
		})

		It("Extracts vendor Color", func() {
			Expect(variant.Color).To(Equal("Schwarz"))
		})

		It("Extracts vendor Price", func() {
			Expect(variant.Price).To(Equal(289.95))
		})
	})

	Context("when no color given", func() {
		BeforeEach(func() {
			exampleFile = "potsdam/nikesuperbad3.0.html"
		})

		It("Extracts title", func() {
			Expect(subject.Name).To(Equal("Nike Super Bad 3.0 schwarz"))
		})

		It("Extracts vendor Size", func() {
			Expect(variant.Size).To(Equal("S"))
		})

		It("Extracts vendor Color", func() {
			Expect(variant.Color).To(Equal("Schwarz"))
		})

		It("Extracts vendor Price", func() {
			Expect(variant.Price).To(Equal(69.95))
		})
	})

})
