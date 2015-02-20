package ganalyse

import (
  "fmt"
  s "strings"
  "strconv"
  "regexp"
  "io"
  "github.com/djimenez/iconv-go"
  "github.com/PuerkitoBio/goquery"
)

type Variant struct {
  id int
  Color, Size, Position string
  Price float64
  Availability int
}

func (v *Variant) String() string {
  // return fmt.Sprintf("id: %d\tcolor: %s\tsize: %s\tpos: %s\tavail: %d\tprice: %.2f", v.id, v.color, v.size, v.position, v.availability, v.price)
  return fmt.Sprintf("color: %s\tsize: %s\tpos: %s\tavail: %d\tprice: %.2f", v.Color, v.Size, v.Position, v.Availability, v.Price)
}

type Product struct{
  id int
  Name string
  Variants []Variant
}

func (p *Product) String() string {
  // out := fmt.Sprintf("id: %d\tname: %s\t variants:", p.id, p.name)
  out := fmt.Sprintf("%s\t variants: %d", p.Name, len(p.Variants))
  // for i := range p.Variants {
  //   out = fmt.Sprintf("%s\n %v", out, p.Variants[i].String())
  // }
  return out
}

func (p *Product) Add(variant Variant) {
  p.Variants = append(p.Variants, variant)
}


// -----------------------------

func Parse(data []byte, charset string) (doc *goquery.Document) {
  var reader io.Reader
  reader = s.NewReader(string(data))
  if charset != "utf-8" {
    reader, _ = iconv.NewReader(reader, charset, "utf-8")
  }
  doc, _ = goquery.NewDocumentFromReader(reader)
  return
}

func NormPrice(value string) (price float64) { //TODO move to model?
  regMatcher := regexp.MustCompile(`([\d,.]+)`)
  r := regMatcher.FindAllStringSubmatch(value, -1)
  if len(r) > 0 {
    value = s.Replace(r[0][1], ",", ".", -1)
    price, _ = strconv.ParseFloat(value, 64)
  }
  return
}
