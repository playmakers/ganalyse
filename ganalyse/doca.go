package ganalyse

import (
  // "fmt"
  s "strings"
  "github.com/PuerkitoBio/goquery"
)

func InspectBossDocA(productPage []byte) Product {
  // sizeMapping := map[string]string {
  //   // "76": "XS",
  //   "1": "S",
  //   "2": "M",
  //   "3": "L",
  //   "4": "XL",
  //   // "31": "XXL",
  //   // "38": "3XL",
  //   // "39": "4XL",
  // }

  reader := s.NewReader(string(productPage))
  doc, _ := goquery.NewDocumentFromReader(reader)

  product := Product {
    name: doc.Find("h1").Text(),
  }

  price := normPrice(doc.Find(".article_details_price strong").Text())

  variant := Variant {
    color: "schwarz",
    size: "L",
    price: price,
    availability: 0,
  }

  product.variants = append(product.variants, variant)

  return product
}

