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
      return map[string]string {
        "7": "Schwarz",
        "10": "Rot",
        "48": "",
        "164": "",
        "51": "",
        "177": "Weiss",
        "217": "navy-blau",
        "223": "silber metallic",
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

