package vendors

import (
  "../ganalyse"
  // "fmt"
  // "github.com/PuerkitoBio/goquery"
)

func InspectDocA(productPage []byte) *ganalyse.Product {
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

  price := ganalyse.NormPrice(doc.Find(".article_details_price strong").Text())

  variant := ganalyse.Variant {
    Color: DEFAULT_COLOR,
    Size: DEFAULT_SIZE,
    Price: price,
    Availability: 0,
  }

  product.Add(variant)

  return &product
}

