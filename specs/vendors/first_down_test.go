package vendors_test

import (
	"github.com/playmakers/ganalyse/lib/vendors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("./Vendors/FirstDown", func() {
	var exampleFile string
	var subject vendors.Product
	var variant vendors.Variant

	JustBeforeEach(func() {
		subject = *vendors.InspectFirstDown(loadExample(exampleFile), "")
		variant = *subject.DefaultVariant()
	})

	BeforeEach(func() {
		exampleFile = "firstdown/bikextremelite.html"
	})

	It("Extracts title", func() {
		Expect(subject.Name).To(Equal("Bike Xtreme Lite 73"))
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

})
