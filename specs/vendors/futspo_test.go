package vendors_test

import (
	"github.com/playmakers/ganalyse/lib/vendors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("./Vendors/Futspo", func() {
	var exampleFile string
	var subject vendors.Product
	var variant vendors.Variant

	JustBeforeEach(func() {
		subject = *vendors.InspectFutspo(loadExample(exampleFile))
		variant = *subject.DefaultVariant()
	})

	BeforeEach(func() {
		exampleFile = "futspo/cutterss60.html"
	})

	It("Extracts title", func() {
		Expect(subject.Name).To(Equal("Gamer S60 LB/ RB Gloves Cutters ( Shockskin Model )"))
	})

	It("Extracts vendor Size", func() {
		Expect(variant.Size).To(Equal("L"))
	})

	It("Extracts vendor Color", func() {
		Expect(variant.Color).To(Equal("charcoal/grey"))
	})

	It("Extracts vendor Price", func() {
		Expect(variant.Price).To(Equal(69.95))
	})

})
