package vendors

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	s "strings"
)

const (
	HOST      = "http://www.dfshop.com/"
	BASE_PATH = "/dfshop/wsDfShop.wsc"
	START_URL = HOST + BASE_PATH + "/pDfStart.p?href=dqbg4p735uvb12kccosjqb6"
	LOGIN     = "PMBE6R47"
)

func login(loginUrl, login string, password string) (requrl string) {
	fmt.Printf("Open: %s\n with: %s \n", loginUrl, password)

	doc, _ := goquery.NewDocument(loginUrl)
	action, _ := doc.Find("#LoginBox form").Attr("action")
	requrl, _ = doc.Find("body").Attr("requrl")

	resp, _ := http.PostForm(HOST+action, url.Values{
		"inpLoginUser": {login},
		"inpLoginPass": {password},
	})
	defer resp.Body.Close()

	return
}

func InspectMeyer(productPage []byte) *Product {
	availabilityMapping := map[string]int{
		"inpQtyRed":    0,
		"inpQtyYellow": 5,
		"inpQtyGreen":  50,
	}

	sizeMapping := map[string]string{
		"2": "M",
		"3": "S",
		"4": "M",
		"5": "L",
		"6": "XL",
		"7": "XXL",
		"8": "3XL",
		"9": "4XL",
	}

	colorMapping := map[string]string{
		"BLK": "Schwarz",
		"BOR": "Orange",
		// "CRD": "",
		// "DGR": "",
		"GLD": "Gelb",
		"GRE": "Grün",
		"KEL": "",
		// "MAR": "",
		"NAV": "Navy-Blau",
		"OGO": "Gold",
		"PUR": "Lila",
		"ROY": "Royal-Blau",
		// "SBG": "",
		"SCA": "Rot",
		"SIL": "Silber",
		// "VGO": "",
		"WHI": "Weiß",
	}

	doc := Parse(productPage, "utf-8")

	productId := func(value string, exists bool) string {
		splitAry := s.Split(value, "=")
		return splitAry[len(splitAry)-1]
	}(doc.Find("meta[property='og:url']").Attr("content"))

	product := &Product{
		Name: doc.Find(fmt.Sprintf("#styledesc%s b", productId)).Text(),
	}

	doc.Find(".tblTrArtRow").Each(func(i int, productSelection *goquery.Selection) {
		color := func(value string) string {
			if len(value) >= 3 {
				if mapping, found := colorMapping[string(value[0:3])]; found {
					return mapping
				}
			}
			return value
		}(productSelection.Find("td b").Text())

		price := NormPrice(productSelection.Next().Find("b").Text())

		productSelection.Next().Find("input[type=text]").Each(func(i2 int, variantSelection *goquery.Selection) {
			size := func(value string, exists bool) string {
				splitAry := s.Split(value, "_")
				if mapping, found := sizeMapping[splitAry[len(splitAry)-1]]; found {
					return mapping
				}
				return value
			}(variantSelection.Attr("name"))

			availability := func(value string, exists bool) int {
				return availabilityMapping[value]
			}(variantSelection.Attr("class"))

			product.AddVariant(size, color, price, availability)
		})
	})

	return product
}
