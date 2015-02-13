package ganalyse

import (
  s "strings"
  "regexp"
  "encoding/csv"
  "github.com/PuerkitoBio/goquery"
)

func InspectFutspo(productPage []byte) Product {
  availabilityMapping := map[string]int {
    "rot": 0,
    "gelb": 5,
    "gruen": 50,
  }

  doc := func(data []byte) *goquery.Document {
    reader := s.NewReader(string(data))
    doc, _ := goquery.NewDocumentFromReader(reader)
    return doc
  }(productPage)

  product := Product {
    name: doc.Find("span.product").Text(),
  }

  variants := doc.Find(".var-ebene script").Last().Text()

  regMatcher := regexp.MustCompile(`new Array\(([^)]+)\)`)
  for _, match := range regMatcher.FindAllStringSubmatch(variants, -1) {
    if len(match) > 1 {
      records := func(csvData string) []string {
        reader := csv.NewReader(s.NewReader(csvData))
        reader.TrimLeadingSpace = true
        result, _ := reader.ReadAll()
        return result[0]
      }(match[1])

      color := func(value string) string {
        return value
      }(records[1])

      price := normPrice(records[4])

      availability := func(value string) int {
        return availabilityMapping[value]
      }(records[6])

      variant := Variant {
        color: color,
        size: records[2],
        price: price,
        availability: availability,
      }

      product.variants = append(product.variants, variant)
    }
  }


  return product
}


