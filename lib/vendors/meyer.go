package vendors

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"regexp"
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

	colorMapping := map[string]string{
		"BLK": "Schwarz",
		"BOR": "Orange",
		"GLD": "Gelb",
		"GRE": "Grün",
		"NAV": "Navy-Blau",
		"OGO": "Gold",
		"PUR": "Lila",
		"ROY": "Royal-Blau",
		"SCA": "Rot",
		"SIL": "Silber",
		"WHI": "Weiß",
		// "KEL": "",
		// "MAR": "",
		// "SBG": "",
		// "VGO": "",
		// "CRD": "",
		// "DGR": "",
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
			regMatcher := regexp.MustCompile(`([\w-]+)`)
	    r := regMatcher.FindAllStringSubmatch(value, -1)
	    if len(r) > 0 {
				if mapping, found := colorMapping[r[0][1]]; found {
					return mapping
				}
			} else {
				value = DEFAULT_COLOR
			}
			return value
		}(productSelection.Find("td b").Text())

		price := func(node *goquery.Selection) float64 {
			priceNode := node.Find("b b")
			if priceNode.Text() == "" {
				priceNode = node.Find("b")
			}
			return NormPrice(priceNode.Text())
		}(productSelection.Next())

		productSelection.Next().Find("input[type=text]").Each(func(i2 int, variantSelection *goquery.Selection) {
			size := func(value string) string {
				return s.TrimSpace(value)
			}(variantSelection.Parent().Text())

			availability := func(value string, exists bool) int {
				return availabilityMapping[value]
			}(variantSelection.Attr("class"))

			product.AddVariant(size, color, price, availability)
		})
	})

	return product
}
