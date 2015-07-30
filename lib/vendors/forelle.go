package vendors

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	s "strings"
)

func InspectForelle(productPage []byte, origin string) *Product {
	colorMapping := map[string]string{
		"1":  "Royal-Blau",
	  "2":  "Orange",
		"3":  "Schwarz",
		"4":  "Weiß",
		"8":  "Navy-Blau",
		"9":  "Rot", // 'Scarlet'
    "12": "Rot", //Cardinal
		"13": "Grün",
		"15": "Gelb",
		"16": "Rot",
		"18": "Lila",
		"46": "Orange",
    "24": "Silber", // "Met. Silver",
    "30": "Met. Gold",
    "31": "Schwarz",
    "32": "Schwarz/Weiß",
    "33": "Royal", // Royal
    "56": "Old Gold",
    "51": "Klar", //Clear, Transparent
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
    "7": "15.0",
    "9": "7.0",
    "10": "8/41",
    "13": "6.5",
    "14": "8.5/42",
    "18": "9/42.5",
    // "9.5/43",
    // "10/44",
    // "10.5/44.5",
    // "11/45",
    "35": "S",
    "33": "M",
    "32": "L",
    "34": "XL",
    "31": "XXL",
    "38": "XXXL",
    "76": "XS",
    UNKNOWN_SIZE: UNKNOWN_SIZE,
    // "",
	}

	doc := Parse(productPage, "utf-8")

	product := &Product{
		Name: doc.Find("h1").First().Text(),
    Origin: origin,
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
    colorFound := false

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
      colorFound = true
		})

    if colorFound == false {
      product.AddVariant(size, DEFAULT_COLOR, price, AVAILABILE)
    }
	}

	return product
}
