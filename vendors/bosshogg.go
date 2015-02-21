package vendors

import (
  "../ganalyse"
  // "fmt"
  // "github.com/PuerkitoBio/goquery"
)

func InspectBossHogg(productPage []byte) *ganalyse.Product {
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

  doc := ganalyse.Parse(productPage, "utf-8")

  product := ganalyse.Product {
    Name: doc.Find("h1").Text(),
  }

  price := ganalyse.NormPrice(doc.Find(".PricesalesPrice").Text())

  variant := ganalyse.Variant {
    Color: DEFAULT_COLOR,
    Size: DEFAULT_SIZE,
    Price: price,
    Availability: 0,
  }

  product.Add(variant)

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

  return &product
}

