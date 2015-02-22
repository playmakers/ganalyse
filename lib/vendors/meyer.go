package vendors

import (
  "../ganalyse"
  "fmt"
  s "strings"
  "github.com/PuerkitoBio/goquery"
  "net/url"
  "net/http"
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

  resp, _ := http.PostForm(HOST + action, url.Values{
    "inpLoginUser": {login},
    "inpLoginPass": {password},
  })
  defer resp.Body.Close()

  return
}

func InspectMeyer(productPage []byte) *ganalyse.Product {
  availabilityMapping := map[string]int {
    "inpQtyRed": 0,
    "inpQtyYellow": 5,
    "inpQtyGreen": 50,
  }

  sizeMapping := map[string]string {
    "3": "S",
    "4": "M",
    "5": "L",
    "6": "XL",
    "7": "XXL",
    "8": "3XL",
    "9": "4XL",
  }

  colorMapping := map[string]string {
    "BLK": "Schwarz",
    "BOR": "",
    "CRD": "",
    "DGR": "",
    "GLD": "",
    "GRE": "",
    "KEL": "",
    "MAR": "",
    "NAV": "",
    "OGO": "",
    "PUR": "",
    "ROY": "",
    "SBG": "",
    "SCA": "",
    "SIL": "",
    "VGO": "",
    "WHI": "",
  }

  doc := ganalyse.Parse(productPage, "utf-8")

  productId := func(value string, exists bool) string {
    splitAry := s.Split(value, "=")
    return splitAry[len(splitAry)-1]
  }(doc.Find("meta[property='og:url']").Attr("content"))

  product := ganalyse.Product {
    Name: doc.Find(fmt.Sprintf("#styledesc%s b", productId)).Text(),
  }

  doc.Find(".tblTrArtRow").Each(func(i int, productSelection *goquery.Selection) {
    color := func(value string) string {
      if len(value) >= 3 {
        return colorMapping[string(value[0:3])]
      } else {
        return ""
      }
    }(productSelection.Find("td b").Text())

    price := ganalyse.NormPrice(productSelection.Next().Find("b").Text())

    productSelection.Next().Find("input[type=text]").Each(func(i2 int, variantSelection *goquery.Selection) {
      size := func(value string, exists bool) string {
        splitAry := s.Split(value, "_")
        return sizeMapping[splitAry[len(splitAry)-1]]
      }(variantSelection.Attr("name"))

      availability := func(value string, exists bool) int {
        return availabilityMapping[value]
      }(variantSelection.Attr("class"))

      product.AddVariant(size, color, price, availability)
    })
  })

  return &product
}
