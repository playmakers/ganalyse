package vendors

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	s "strings"
)

type Variant struct {
	Color, Size, Position string
	Price                 float64
	Availability          int
}

func (v *Variant) String() string {
	// return fmt.Sprintf("id: %d\tcolor: %s\tsize: %s\tpos: %s\tavail: %d\tprice: %.2f", v.id, v.color, v.size, v.position, v.availability, v.price)
	return fmt.Sprintf("color: %s\tsize: %s\tpos: %s\tavail: %d\tprice: %.2f", v.Color, v.Size, v.Position, v.Availability, v.Price)
}

type Product struct {
	Name     string
	Origin  string
	Variants map[string]*Variant
}

func (p *Product) String() (out string) {
	return fmt.Sprintf("%s\t variant: %v", p.Name, p.DefaultVariant().String())
}

func (p *Product) key(size string, color string) string {
	return fmt.Sprintf("%s-%s", s.ToLower(size), s.ToLower(color))
}

// func (p *Product) UpdateVariant(color string, size string, position string, availability int) {
//   for _, variant := range p.Variants {
//     // fmt.Printf("test: %s %s || %s %s\n", color, variant.Color, size, variant.Size)
//     if color == variant.Color && size == variant.Size {
//       variant.Availability = availability
//       fmt.Printf(" --> match: %s %s || %s %s\n", color, variant.Color, size, variant.Size)
//       // fmt.Printf("match: %d\n", variant)
//     }
//   }
// }

func (p *Product) DefaultVariant() (variant *Variant) {
	variant = p.Get(p.key("L", "Schwarz"))
	if variant == nil {
		sizes := p.findAll(func(v *Variant) string { return v.Size })
		colors := p.findAll(func(v *Variant) string { return v.Color })
		size := p.findOrSmallest("L", func(v *Variant) string { return v.Size })
		color := p.findOrSmallest("Schwarz", func(v *Variant) string { return v.Color })
		log.Println(" No default found in:", sizes, colors, " using instead:", size, color)
		variant = p.Get(p.key(size, color))
	}
	return
}

type attrResolver func(*Variant) string

func (p *Product) findOrSmallest(token string, attr attrResolver) string {
	smallest := "zz"
	for _, variant := range p.Variants {
		found := attr(variant)
		if token == found {
			return found
		}
		if found < smallest {
			smallest = found
		}
	}
	return smallest
}

func (p *Product) findAll(attr attrResolver) (keys []string) {
	values := map[string]int{}
	for _, variant := range p.Variants {
		values[attr(variant)] = 0
	}

	for key, _ := range values {
		keys = append(keys, key)
	}
	return
}

func (p *Product) AddVariant(size string, color string, price float64, availability int) {
	if p.Variants == nil {
		p.Variants = make(map[string]*Variant)
	}

	p.Variants[p.key(size, color)] = &Variant{
		Size:         size,
		Color:        color,
		Price:        price,
		Availability: availability,
	}
}

func (p *Product) Get(key string) *Variant {
	if p.Variants != nil {
		if val, ok := p.Variants[key]; ok {
			return val
		}
	}
	return nil
}

// -----------------------------

func NormPrice(value string) (price float64) { //TODO move to model?
	// log.Println(value)
	regMatcher := regexp.MustCompile(`([\d,.]+)`)
	r := regMatcher.FindAllStringSubmatch(value, -1)
	// log.Println(r)
	if len(r) > 0 {
		value = s.Replace(r[0][1], ",", ".", -1)
		price, _ = strconv.ParseFloat(value, 64)
	}
	return
}
