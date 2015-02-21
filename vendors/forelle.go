package vendors

import (
  "../ganalyse"
)

func InspectForelle(productPage []byte) *ganalyse.Product {
  doc := ganalyse.Parse(productPage, "utf-8")

  product := ganalyse.Product {
    Name: doc.Find("h1").First().Text(),
  }

  price := ganalyse.NormPrice(doc.Find(".art-price").Text())

  sizes := getValues(
    doc.Find(".sizes input"),
    DEFAULT_SIZE,
    func(value string) string {
      return map[string]string {
        "76": "XS",
        "35": "S",
        "33": "M",
        "32": "L",
        "34": "XL",
        "31": "XXL",
        "38": "3XL",
      }[value]
    },
  )

  colors := getValues(
    doc.Find(".colors input"),
    DEFAULT_COLOR,
    func(value string) string {
      return value
      return map[string]string {
        "1": "Royal",
        "2": "Orange",
        "4": "White",
        "8": "Navy",
        "13": "Forest",
        "15": "Yellow",
        "16": "Maroon",
        "18": "Purple",
      }[value]
    },
  )

  for _, size := range sizes {
    for _, color := range colors {
      product.AddVariant(size, color, price, DEFAULT_AVAILABILITY)
    }
  }

  return &product
}
