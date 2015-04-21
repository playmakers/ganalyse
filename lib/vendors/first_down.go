package vendors

// TODO
import (
	"github.com/playmakers/ganalyse/lib/ganalyse"
	"regexp"
	// "github.com/PuerkitoBio/goquery"
)

func InspectFirstDown(productPage []byte) *ganalyse.Product {
	doc := ganalyse.Parse(productPage, "iso-8859-1")

	product := &ganalyse.Product{
		Name: doc.Find("h1").Text(),
	}

	price := ganalyse.NormPrice(doc.Find("#aprice").Text())

	sizes := getSizes(
		findOption(doc.Find("select"), "Größe: "),
		regexp.MustCompile(`(S|M|L|X?X?XL|\dX)[^+]*(\+([\d,.]+))?`),
	)
	colors := getColors(
		dropFirst(findOption(doc.Find("select"), "Farbe: ")),
	)

	for _, sizeAndPrice := range sizes {
		for _, color := range colors {
			product.AddVariant(sizeAndPrice.size, color, price+sizeAndPrice.price, DEFAULT_AVAILABILITY)
		}
	}

	return product
}
