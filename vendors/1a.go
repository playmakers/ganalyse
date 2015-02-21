package vendors

import (
  // "fmt"
  "../ganalyse"
  s "strings"
  "regexp"
  "github.com/PuerkitoBio/goquery"
)

type sizeWithPrice struct {
  size string
  price float64
}

type mapping func(string) string

const (
  DEFAULT_SIZE  = "l"
  DEFAULT_COLOR = "schwarz"
)
func getValues(selection *goquery.Selection, defaultValue string, mapping mapping) (values []string) {
  selection.Each(func(i int, valueSelection *goquery.Selection) {
    value := func(value string, exists bool) string {
      return mapping(s.ToLower(s.TrimSpace(value)))
    }(valueSelection.Attr("value"))

    values = append(values, value)
  })
  if len(values) < 1 {
    values = append(values, defaultValue)
  }
  return
}

func findOption(haystack *goquery.Selection, needle string) *goquery.Selection {
  return haystack.FilterFunction(func(i int, selection *goquery.Selection) bool {
    return selection.Parent().Prev().Text() == needle
  }).Find("option")
}

func getSizes(selection *goquery.Selection, regMatcher *regexp.Regexp) (sizes []sizeWithPrice) {
  selection.Each(func(i int, sizeSelection *goquery.Selection) {
    sizeString := sizeSelection.Text()
    r := regMatcher.FindAllStringSubmatch(sizeString, -1)
    if len(r) > 0 {
      sizes = append(sizes, sizeWithPrice {
        size: s.ToLower(r[0][1]),
        price: ganalyse.NormPrice(r[0][3]),
      })
    }
  })
  if(len(sizes) < 1) {
    sizes = append(sizes, sizeWithPrice {
      size: DEFAULT_SIZE,
      price: 0,
    })
  }
  return
}

func getColors(selection *goquery.Selection) (colors []string) {
  selection.Each(func(i int, colorSelection *goquery.Selection) {
    if i > 0 {
      colors = append(colors, func(value string) string {
        return s.ToLower(s.TrimSpace(value))
      }(colorSelection.Text()))
    }
  })
  if(len(colors) < 1) {
    colors = append(colors, DEFAULT_COLOR)
  }
  return
}

func Inspect1A(productPage []byte) *ganalyse.Product {
  doc := ganalyse.Parse(productPage, "iso-8859-1")

  product := ganalyse.Product {
    Name: doc.Find("h1").Text(),
  }

  price := ganalyse.NormPrice(doc.Find("#price").Text())

  sizes  := getSizes(
    findOption(doc.Find("select"), "Größe"),
    regexp.MustCompile(`(S|M|L|\d?X?XL)[^+]*(\+ ([\d,]+))?`),
  )
  colors := getColors(
    findOption(doc.Find("select"), "Farbe"),
  )

  for _, sizeAndPrice := range sizes {
    for _, color := range colors {

      product.Add(ganalyse.Variant {
        Color: color,
        Size: sizeAndPrice.size,
        Price: price + sizeAndPrice.price,
        Availability: 0,
      })
    }
  }

  return &product
}

