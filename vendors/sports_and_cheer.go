package vendors

import (
  // "fmt"
  "../ganalyse"
  s "strings"
  "github.com/PuerkitoBio/goquery"
)

func InspectSportsAndCheer(productPage []byte) *ganalyse.Product {
  doc := ganalyse.Parse(productPage, "iso-8859-1")

  product := ganalyse.Product {
    Name: s.TrimSpace(doc.Find("h1").Text()),
  }

  price := func(value string, exists bool) float64 {
    return ganalyse.NormPrice(value)
  }(doc.Find("input[name=\"vk_brutto\"]").Attr("value"))

  doc.Find("select[name=\"a_groesse\"] option").Each(func(i int, sizeSelection *goquery.Selection) { // TODO loop at least once!
    if(i < 2) { return }

    size := func(value string, exists bool) string {
        return s.TrimSpace(value)
    }(sizeSelection.Attr("value"))

    doc.Find("select[name=\"a_farbe\"] option").Each(func(i int, colorSelection *goquery.Selection) {
      if(i < 2) { return }

      color :=func(value string, exists bool) string {
        return s.TrimSpace(value)
      }(colorSelection.Attr("value"))

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

