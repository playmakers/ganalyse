package vendors

import (
  "../ganalyse"
  "github.com/PuerkitoBio/goquery"
)

func InspectPotsdam(productPage []byte) *ganalyse.Product {
  sizeMapping := map[string]string {
    // "76": "XS",
    "1": "S",
    "2": "M",
    "3": "L",
    "4": "XL",
    // "31": "XXL",
    // "38": "3XL",
    // "39": "4XL",
  }

  doc := ganalyse.Parse(productPage, "utf-8")

  product := ganalyse.Product {
    Name: doc.Find(".productName").Text(),
  }

  price := ganalyse.NormPrice(doc.Find(".productPrice").Text())

  doc.Find("select[name=\"id[2]\"] option").Each(func(i int, sizeSelection *goquery.Selection) { // TODO loop at least once!
    size := func(value string, exists bool) string {
        return sizeMapping[value] // TODO add to price
    }(sizeSelection.Attr("value"))

    doc.Find("select[name=\"id[1]\"] option").Each(func(i int, colorSelection *goquery.Selection) {
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

