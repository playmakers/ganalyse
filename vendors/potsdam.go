package vendors

import (
  "../ganalyse"
)

func InspectPotsdam(productPage []byte) *ganalyse.Product {
  doc := ganalyse.Parse(productPage, "utf-8")

  product := ganalyse.Product {
    Name: doc.Find(".productName").Text(),
  }

  price := ganalyse.NormPrice(doc.Find(".productPrice").Text())

  sizes := getValues(
    doc.Find("select[name=\"id[2]\"] option"),
    DEFAULT_SIZE,
    func(value string) string {
      return map[string]string {
        "1": "S",
        "2": "M",
        "3": "L",
        "4": "XL",
      }[value]
    },
  )

  colors := getValues(
    doc.Find("select[name=\"id[1]\"] option"),
    DEFAULT_COLOR,
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

