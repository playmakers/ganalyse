package vendors

import (
  s "strings"
  "../ganalyse"
  "github.com/PuerkitoBio/goquery"
)

func InspectUsSportshop(productPage []byte) *ganalyse.Product {
  sizeMapping := map[string]string {
    "14": "S",
    "15": "M",
    "16": "L",
    "17": "XL",
    "18": "XXL",
    "19": "3XL",
    "40": "4XL",
  }

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

  doc.Find("select[name=\"id[2]\"] option").Each(func(i int, sizeSelection *goquery.Selection) { // TODO loop at least once!
    size := func(value string, exists bool) string {
        return sizeMapping[value]
    }(sizeSelection.Attr("value"))

    doc.Find("select[name=\"id[1]\"] option").Each(func(i int, colorSelection *goquery.Selection) {
      color, _ := colorSelection.Attr("value") //TODO color mapping

      variant := ganalyse.Variant {
        Color: color,
        Size: size,
        Price: price,
        Availability: 0,
      }

      product.Add(variant)
    })
  })

  return &product
}

