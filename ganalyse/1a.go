package ganalyse

import (
  // "fmt"
  s "strings"
  "regexp"
  "github.com/djimenez/iconv-go"
  "github.com/PuerkitoBio/goquery"
)

func findOption(haystack *goquery.Selection, needle string) *goquery.Selection {
  return haystack.FilterFunction(func(i int, selection *goquery.Selection) bool {
    return selection.Parent().Prev().Text() == needle
  }).Find("option")
}

func Inspect1A(productPage []byte) Product {
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

  doc := func(data []byte) *goquery.Document {
    reader, _ := iconv.NewReader(s.NewReader(string(data)), "iso-8859-1", "utf-8")
    doc, _ := goquery.NewDocumentFromReader(reader)
    return doc
  }(productPage)

  product := Product {
    name: doc.Find("h1").Text(),
  }

  price := normPrice(doc.Find("#price").Text())

  findOption(doc.Find("select"), "Größe").Each(func(i int, sizeSelection *goquery.Selection) {
    size, extraPrice := func(value string) (size string, extraPrice float64) {
      regMatcher := regexp.MustCompile(`(S|M|L|\d?X?XL)[^+]*(\+ ([\d,]+))?`)
      r := regMatcher.FindAllStringSubmatch(value, -1)
      if len(r) > 0 {
        size = r[0][1]
        extraPrice = normPrice(r[0][3])
      }
      return
    }(sizeSelection.Text())

    if len(size) > 0 {
      findOption(doc.Find("select"), "Farbe").Each(func(i int, colorSelection *goquery.Selection) {
        if i > 0 {
          color := func(value string) string {
            return s.TrimSpace(value)
          }(colorSelection.Text())

          variant := Variant {
            color: color,
            size: size,
            price: price + extraPrice,
            availability: 0,
          }

          product.variants = append(product.variants, variant)
        }
      })
    }
  })

  return product
}

