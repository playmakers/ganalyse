package vendors

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	s "strings"
)

func InspectForelle(productPage []byte) *Product {
	colorMapping := map[string]string{
		"1":  "Royal-Blau",
	  "2":  "Orange",
		"3":  "Schwarz",
		"4":  "Weiß",
		"8":  "Navy-Blau",
		"9":  "Rot", // 'Scarlet'
		"13": "Grün",
		"15": "Gelb",
		"16": "Rot",
		"18": "Lila",
		"46": "Orange",
	}

    // 'Royal'           => '',
    // 'Black'           => 'Schwarz',
    // 'Black/Black'     => 'Schwarz',
    // 'Black/White'     => '',
    // 'Cardinal'        => 'Rot',
    // 'Clear'           => nil,
    // 'Maroon'          => 'Rot',
    // 'Scarlet'         => 'Rot',
    // 'Forest'          => 'Grün',
    // 'Met. Silver'     => 'Silber',
    // 'Navy'            => 'Navy-Blau',
    // 'Met. Gold'       => 'Gold',
    // 'Notre Dame Gold' => 'Gold',
    // 'Old Gold'        => 'Gold',
    // 'Vegas Gold'      => 'Gold',
    // 'Purple'          => 'Lila',
    // 'White'           => 'Weiß',
    // 'Yellow'          => 'Gelb',
    // 'Orange'          => '',
    // 'Pink'            => '',

	sizeMapping := map[string]string{
		"76": "XS",
		"35": "S",
		"33": "M",
		"32": "L",
		"34": "XL",
		"31": "XXL",
		"38": "3XL",
	}

	doc := Parse(productPage, "utf-8")

	product := &Product{
		Name: doc.Find("h1").First().Text(),
	}

	price := NormPrice(doc.Find(".art-price").Text())

	sizes := getValues(
		doc.Find(".sizes input"),
		UNKNOWN_SIZE,
		func(value string) string {
			return value
		},
	)

	for _, sizeValue := range sizes {
		size := sizeMapping[sizeValue]

		selector := func() string {
			if sizeValue == UNKNOWN_SIZE {
				return ".colors .item input"
			} else {
				return fmt.Sprintf(".colors .option_item_%s input", sizeValue)
			}
		}()

		doc.Find(selector).Each(func(i int, colorSelection *goquery.Selection) {
			color := func(value string, exists bool) string {
				mapping, ok := colorMapping[s.TrimSpace(value)]
				if !ok {
					return value
				}
				return mapping
			}(colorSelection.Attr("value"))

			availability := func(value string, exists bool) int {
				if(s.Contains(value, "Out of stock")) {
					return OUTOFSTOCK
				} else {
					return AVAILABILE
				}
			}(colorSelection.Parent().Attr("title"))

			product.AddVariant(size, color, price, availability)
		})

	}

	return product
}
