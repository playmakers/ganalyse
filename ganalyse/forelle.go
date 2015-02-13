package ganalyse

import (
  s "strings"
  "github.com/PuerkitoBio/goquery"
)

func InspectForelle(productPage []byte) Product {
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

  reader := s.NewReader(string(productPage))
  doc, _ := goquery.NewDocumentFromReader(reader)

  product := Product {
    name: doc.Find("h1").Text(),
  }

  price := normPrice(doc.Find(".art-price").Text())

  doc.Find(".sizes input").Each(func(i int, sizeSelection *goquery.Selection) {
    size := func(value string, exists bool) string {
        return sizeMapping[value]
    }(sizeSelection.Attr("value"))

    doc.Find(".colors input").Each(func(i int, colorSelection *goquery.Selection) {
      color, _ := colorSelection.Attr("value")

      variant := Variant {
        color: color,
        size: size,
        price: price,
        availability: 0,
      }

      product.variants = append(product.variants, variant)
    })
  })

  return product
}

