package vendors

import (
  "../ganalyse"
  s "strings"
)

func InspectSportsAndCheer(productPage []byte) *ganalyse.Product {
  doc := ganalyse.Parse(productPage, "iso-8859-1")

  product := &ganalyse.Product {
    Name: s.TrimSpace(doc.Find("h1").Text()),
  }

  price := func(value string, exists bool) float64 {
    return ganalyse.NormPrice(value)
  }(doc.Find("input[name=\"vk_brutto\"]").Attr("value"))

  sizes := getValues(
    dropFirst(doc.Find("select[name=\"a_groesse\"] option")),
    DEFAULT_SIZE,
    func(value string) string {
      return value
    },
  )

  colors := getValues(
    dropFirst(doc.Find("select[name=\"a_farbe\"] option")),
    DEFAULT_COLOR,
    func(value string) string {
      return value
    },
  )

  for _, size := range sizes {
    for _, color := range colors {
      product.AddVariant(size, color, price, DEFAULT_AVAILABILITY)
    }
  }

  return product
}

