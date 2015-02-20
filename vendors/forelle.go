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
    "L",
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
    "schwarz",
    func(value string) string {
      return value
    },
  )

  for _, size := range sizes {
    for _, color := range colors {

      product.Add(ganalyse.Variant {
        Color: color,
        Size: size,
        Price: price,
        Availability: 0,
      })
    }
  }

  return &product
}
