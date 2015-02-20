package vendors

// TODO
import (
  "../ganalyse"
  "regexp"
  // "github.com/PuerkitoBio/goquery"
)

func InspectFirstDown(productPage []byte) *ganalyse.Product {
  doc := ganalyse.Parse(productPage, "iso-8859-1")

  product := ganalyse.Product {
    Name: doc.Find("h1").Text(),
  }

  price := ganalyse.NormPrice(doc.Find("#aprice").Text())

  sizes  := getSizes(
    findOption(doc.Find("select"), "Größe: "),
    regexp.MustCompile(`(S|M|L|X?X?XL|\dX)[^+]*(\+([\d,.]+))?`),
  )
  colors := getColors(
    findOption(doc.Find("select"), "Farbe: "),
  )

  for _, sizeAndPrice := range sizes {
    for _, color := range colors {

      product.Add(ganalyse.Variant {
        Color: color,
        Size: sizeAndPrice.size,
        Price: price + sizeAndPrice.price,
        Availability: 0,
      })
    }
  }

  return &product
}

