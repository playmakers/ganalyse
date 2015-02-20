package vendors

import (
  "../ganalyse"
  s "strings"
  "github.com/PuerkitoBio/goquery"
)

func dropFirst(selection *goquery.Selection) *goquery.Selection {
  if selection.Size() > 0 {
    return selection.Slice(1, -1)
  } else {
    return selection.Remove()
  }
}

func InspectSportsAndCheer(productPage []byte) *ganalyse.Product {
  doc := ganalyse.Parse(productPage, "iso-8859-1")

  product := ganalyse.Product {
    Name: s.TrimSpace(doc.Find("h1").Text()),
  }

  price := func(value string, exists bool) float64 {
    return ganalyse.NormPrice(value)
  }(doc.Find("input[name=\"vk_brutto\"]").Attr("value"))

  sizes := getValues(
    dropFirst(doc.Find("select[name=\"a_groesse\"] option")),
    "L",
    func(value string) string {
      return value
    },
  )

  colors := getValues(
    dropFirst(doc.Find("select[name=\"a_farbe\"] option")),
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

