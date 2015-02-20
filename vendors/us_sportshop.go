package vendors

import (
  s "strings"
  "../ganalyse"
)

func InspectUsSportshop(productPage []byte) *ganalyse.Product {
  doc := ganalyse.Parse(productPage, "iso-8859-1")

  name := func(value string, e error) string {
    splitAry := s.Split(value, "<br/>")
    return splitAry[0]
  }(doc.Find("h1").Last().Html())

  product := ganalyse.Product {
    Name: name,
  }

  price := func(value string, e error) float64 {
    splitAry := s.Split(value, "<br/>")
    return ganalyse.NormPrice(splitAry[0])
  }(doc.Find("h1").First().Html())

  sizes := getValues(
    doc.Find("select[name=\"id[2]\"] option"),
    "L",
    func(value string) string {
      return map[string]string {
        "14": "S",
        "15": "M",
        "16": "L",
        "17": "XL",
        "18": "XXL",
        "19": "3XL",
        "40": "4XL",
      }[value]
    },
  )

  colors := getValues(
    doc.Find("select[name=\"id[1]\"] option"),
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

