package ganalyse

import (
  "fmt"
  s "strings"
  "strconv"
  "regexp"
)

type Variant struct {
  id int
  color, size, position string
  price float64
  availability int
}

func (v *Variant) String() string {
  // return fmt.Sprintf("id: %d\tcolor: %s\tsize: %s\tpos: %s\tavail: %d\tprice: %.2f", v.id, v.color, v.size, v.position, v.availability, v.price)
  return fmt.Sprintf("color: %s\tsize: %s\tpos: %s\tavail: %d\tprice: %.2f", v.color, v.size, v.position, v.availability, v.price)
}

type Product struct{
  id int
  name string
  variants []Variant
}

func (p *Product) String() string {
  // out := fmt.Sprintf("id: %d\tname: %s\t variants:", p.id, p.name)
  out := fmt.Sprintf("%s\t variants:", p.name)
  for i := range p.variants {
    out = fmt.Sprintf("%s\n %v", out, p.variants[i].String())
  }
  return out
}

func normPrice(value string) (price float64) { //TODO move to model?
  regMatcher := regexp.MustCompile(`([\d,.]+)`)
  r := regMatcher.FindAllStringSubmatch(value, -1)
  if len(r) > 0 {
    value = s.Replace(r[0][1], ",", ".", -1)
    price, _ = strconv.ParseFloat(value, 32)
  }
  return
}
