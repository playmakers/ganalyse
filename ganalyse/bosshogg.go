package ganalyse

import (
  // "fmt"
  s "strings"
  "github.com/PuerkitoBio/goquery"
)

func InspectBossHogg(productPage []byte) Product {
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

  price := normPrice(doc.Find(".PricesalesPrice").Text())

  variant := Variant {
    color: "schwarz",
    size: "L",
    price: price,
    availability: 0,
  }

  product.variants = append(product.variants, variant)

  // doc.Find("select option").Each(func(i int, sizeSelection *goquery.Selection) { // TODO loop at least once!
  //   // size := func(value string, exists bool) string {
  //   //     return sizeMapping[value] // TODO add to price
  //   // }(sizeSelection.Attr("value"))

  //   // doc.Find("select[name=\"id[1]\"] option").Each(func(i int, colorSelection *goquery.Selection) {
  //   //   color, _ := colorSelection.Attr("value")

  //     variant := Variant {
  //       color: color,
  //       size: size,
  //       price: price,
  //       availability: 0,
  //     }

  //     product.variants = append(product.variants, variant)
  //   })
  // })

  return product
}

