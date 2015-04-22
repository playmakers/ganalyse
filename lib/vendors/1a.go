package vendors

import (
	"regexp"
)

func Inspect1A(productPage []byte) *Product {
	doc := Parse(productPage, "iso-8859-1")

	product := &Product{
		Name: doc.Find("h1").Text(),
	}

	price := NormPrice(doc.Find("#price").Text())

	sizes := getSizes(
		findOption(doc.Find("select"), "Größe"),
		regexp.MustCompile(`(S|M|L|\d?X?XL)[^+]*(\+ ([\d,]+))?`),
	)
	colors := getColors(
		findOption(doc.Find("select"), "Farbe"),
	)

	for _, sizeAndPrice := range sizes {
		for _, color := range colors {
			product.AddVariant(sizeAndPrice.size, color, price+sizeAndPrice.price, DEFAULT_AVAILABILITY)
		}
	}

	return product
}
