package vendors

import (
  // "fmt"
  "../ganalyse"
  s "strings"
  "regexp"
  "github.com/PuerkitoBio/goquery"
)

func findOption(haystack *goquery.Selection, needle string) *goquery.Selection {
  return haystack.FilterFunction(func(i int, selection *goquery.Selection) bool {
    return selection.Parent().Prev().Text() == needle
  }).Find("option")
}

func Inspect1A(productPage []byte) *ganalyse.Product {
  // sizeMapping := map[string]string {
  //   // "XS":  "XS",
  //   "S":   "S",
  //   "M":   "M",
  //   "L":   "L",
  //   "XL":  "XL",
  //   "XXL": "XXL",
  //   "3XL": "3XL",
  //   // "39": "4XL",
  // }

  doc := ganalyse.Parse(productPage, "iso-8859-1")

  product := ganalyse.Product {
    Name: doc.Find("h1").Text(),
  }

  price := ganalyse.NormPrice(doc.Find("#price").Text())

  findOption(doc.Find("select"), "Größe").Each(func(i int, sizeSelection *goquery.Selection) {
    size, extraPrice := func(value string) (size string, extraPrice float64) {
      regMatcher := regexp.MustCompile(`(S|M|L|\d?X?XL)[^+]*(\+ ([\d,]+))?`)
      r := regMatcher.FindAllStringSubmatch(value, -1)
      if len(r) > 0 {
        size = r[0][1]
        extraPrice = ganalyse.NormPrice(r[0][3])
      }
      return
    }(sizeSelection.Text())

    if len(size) > 0 {
      findOption(doc.Find("select"), "Farbe").Each(func(i int, colorSelection *goquery.Selection) {
        if i > 0 {
          color := func(value string) string {
            return s.TrimSpace(value)
          }(colorSelection.Text())

          variant := ganalyse.Variant {
            Color: color,
            Size: size,
            Price: price + extraPrice,
            Availability: 0,
          }

          product.Add(variant)
        }
      })
    }
  })

  return &product
}

