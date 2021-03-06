package vendors

import (
	"regexp"
)

func InspectFirstDown(productPage []byte, origin string) *Product {
	doc := Parse(productPage, "iso-8859-1")

	product := &Product{
		Name: doc.Find("h1").Text(),
		Origin: origin,
	}

	price := NormPrice(doc.Find("#aprice").Text())

	sizes := getSizes(
		findOption(doc.Find("select"), "Größe: "),
		regexp.MustCompile(`(S|M|L|X?X?XL|\dX)[^+]*(\+([\d,.]+))?`),
	)
	colors := getColors(
		dropFirst(findOption(doc.Find("select"), "Farbe: ")),
	)

	for _, sizeAndPrice := range sizes {
		for _, color := range colors {
			product.AddVariant(sizeAndPrice.size, color, price+sizeAndPrice.price, AVAILABILE)
		}
	}

	return product
}
