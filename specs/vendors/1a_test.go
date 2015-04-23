package vendors_test

import (
	"github.com/playmakers/ganalyse/lib/vendors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("./Vendors/1a", func() {
	var exampleFile string
	var subject vendors.Product
	var variant vendors.Variant

	JustBeforeEach(func() {
		subject = *vendors.Inspect1A(loadExample(exampleFile))
		variant = *subject.DefaultVariant()
	})

	Context("when no Variants are available", func() {
		BeforeEach(func() {
			exampleFile = "1a/bikextremelite.html"
		})

		It("Extracts title", func() {
			Expect(subject.Name).To(Equal("X-treme Lite RB/DB"))
		})

		It("Extracts vendor Size", func() {
			Expect(variant.Size).To(Equal("L"))
		})

		It("Extracts vendor Color", func() {
			Expect(variant.Color).To(Equal("Schwarz"))
		})

		It("Extracts vendor Price", func() {
			Expect(variant.Price).To(Equal(169.5))
		})
	})

	Context("when Variants are available", func() {
		BeforeEach(func() {
			exampleFile = "1a/cutterss60.html"
		})

		It("Extracts vendor Size", func() {
			Expect(variant.Size).To(Equal("L"))
		})

		It("Extracts vendor Color", func() {
			Expect(variant.Color).To(Equal("Schwarz"))
		})

		It("Extracts vendor Price", func() {
			Expect(variant.Price).To(Equal(57.5))
		})
	})

})
