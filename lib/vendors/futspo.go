package vendors

import (
  "../ganalyse"
  s "strings"
  "regexp"
  "encoding/csv"
)

func InspectFutspo(productPage []byte) *ganalyse.Product {
  availabilityMapping := map[string]int {
    "rot": 0,
    "gelb": 5,
    "gruen": 50,
  }

  doc := ganalyse.Parse(productPage, "iso-8859-1")

  if doc.Find("title").Text() == "futspo.de - Bitte beachten Sie!" {
    return nil
  }

  product := &ganalyse.Product {
    Name: doc.Find("span.product").Text(),
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

      price := ganalyse.NormPrice(records[4])

      availability := func(value string) int {
        return availabilityMapping[value]
      }(records[6])

      product.AddVariant(records[2], color, price, availability)
    }
  }


  return product
}


