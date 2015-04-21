package vendors

import (
	"github.com/playmakers/ganalyse/lib/ganalyse"
	"regexp"
)

func Inspect1A(productPage []byte) *ganalyse.Product {
	doc := ganalyse.Parse(productPage, "iso-8859-1")

	product := &ganalyse.Product{
		Name: doc.Find("h1").Text(),
	}

	price := ganalyse.NormPrice(doc.Find("#price").Text())

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
