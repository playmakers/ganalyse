package vendors

import (
	"github.com/playmakers/ganalyse/lib/ganalyse"
	s "strings"
)

func InspectUsSportshop(productPage []byte) *ganalyse.Product {
	doc := ganalyse.Parse(productPage, "iso-8859-1")

	name := func(value string, e error) string {
		splitAry := s.Split(value, "<br/>")
		return splitAry[0]
	}(doc.Find("h1").Last().Html())

	product := &ganalyse.Product{
		Name: name,
	}

	price := func(value string, e error) float64 {
		splitAry := s.Split(value, "<br/>")
		return ganalyse.NormPrice(splitAry[0])
	}(doc.Find("h1").First().Html())

	sizes := getValues(
		doc.Find("select[name=\"id[2]\"] option"),
		DEFAULT_SIZE,
		func(value string) string {
			return map[string]string{
				"14": "S",
				"15": "M",
				"16": "L",
				"17": "XL",
				"18": "XXL",
				"19": "3XL",
				"40": "4XL",
			}[value]
		},
	)

	colors := getValues(
		doc.Find("select[name=\"id[1]\"] option"),
		DEFAULT_COLOR,
		func(value string) string {
			return map[string]string{
				"GLD": "",
				"SCA": "",
				"BOR": "",
				"NAV": "",
				"MAR": "",
				"PUR": "",
				"GRE": "",
				"BLK": "Schwarz",
				"ROY": "",
				"SIL": "",
				"DGR": "",
				"OGO": "",
				"WHI": "",
				"KEL": "",
				"VGO": "",
				"CRD": "",
				"SBG": "",
				"24":  "Navy",
			}[value]
		},
	)

	for _, size := range sizes {
		for _, color := range colors {
			product.AddVariant(size, color, price, DEFAULT_AVAILABILITY)
		}
	}

	return product
}
