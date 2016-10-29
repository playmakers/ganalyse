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
	var sndVariant vendors.Variant

	JustBeforeEach(func() {
		subject = *vendors.InspectMeyer(loadExample(exampleFile), "")
		variant = *subject.DefaultVariant()
	})

	Context("xenithepic", func() {
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

	Context("bikextremelite", func() {
		BeforeEach(func() {
			exampleFile = "meyer/bikextremelite.html"
		})

		It("Extracts title", func() {
			Expect(subject.Name).To(Equal("BASH72   BIKE Xtreme Lite OL/DL Pad"))
		})

		It("Extracts vendor Size", func() {
			Expect(variant.Size).To(Equal("L"))
		})

		It("Extracts vendor Color", func() {
			Expect(variant.Color).To(Equal("Schwarz"))
		})

		It("Extracts vendor Price", func() {
			Expect(variant.Price).To(Equal(180.0))
		})

		Context("second variant", func() {
			BeforeEach(func() {
				sndVariant = *subject.Variants["xl-schwarz"]
			})

			It("Extracts vendor Size", func() {
				Expect(sndVariant.Size).To(Equal("XL"))
			})
		})
	})

})
