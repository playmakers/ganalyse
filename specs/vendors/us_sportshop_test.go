package vendors_test

import (
	"github.com/playmakers/ganalyse/lib/vendors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("./Vendors/UsSportshop", func() {
	var exampleFile string
	var subject vendors.Product
	var variant vendors.Variant

	JustBeforeEach(func() {
		subject = *vendors.InspectUsSportshop(loadExample(exampleFile), "")
		variant = *subject.DefaultVariant()
	})

	BeforeEach(func() {
		exampleFile = "ussportshop/nikesuperbad3.0.html"
	})

	It("Extracts title", func() {
		Expect(subject.Name).To(Equal("Nike Superbad 3 Padded-Receiver Glove - navy"))
	})

	It("Extracts vendor Size", func() {
		Expect(variant.Size).To(Equal("L"))
	})

	It("Extracts vendor Color", func() {
		Expect(variant.Color).To(Equal("Navy"))
	})

	It("Extracts vendor Price", func() {
		Expect(variant.Price).To(Equal(69.9))
	})

})
