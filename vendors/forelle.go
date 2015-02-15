package vendors

import (
  "../ganalyse"
  "github.com/PuerkitoBio/goquery"
)

func InspectForelle(productPage []byte) *ganalyse.Product {
  sizeMapping := map[string]string {
    "76": "XS",
    "35": "S",
    "33": "M",
    "32": "L",
    "34": "XL",
    "31": "XXL",
    "38": "3XL",
    // "39": "4XL",
  }

  doc := ganalyse.Parse(productPage, "utf-8")

  product := ganalyse.Product {
    Name: doc.Find("h1").First().Text(),
  }

  price := ganalyse.NormPrice(doc.Find(".art-price").Text())

  doc.Find(".sizes input").Each(func(i int, sizeSelection *goquery.Selection) {
    size := func(value string, exists bool) string {
        return sizeMapping[value]
    }(sizeSelection.Attr("value"))

    doc.Find(".colors input").Each(func(i int, colorSelection *goquery.Selection) {
      color, _ := colorSelection.Attr("value")

      variant := ganalyse.Variant {
        Color: color,
        Size: size,
        Price: price,
        Availability: 0,
      }

      product.Add(variant)
    })
  })

  return &product
}

