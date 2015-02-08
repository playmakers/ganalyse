// curl is a simple cURL replacement.
package main

import (
  "fmt"
  "log"
  "os"
  "io/ioutil"
  s "strings"
  "strconv"
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

type Variant struct {
  id int
  color, size, position string
  price float64
  availability int
}

func (v *Variant) String() string {
  return fmt.Sprintf("id: %d\tcolor: %s\tsize: %s\tpos: %s\tavail: %d\tprice: %.2f", v.id, v.color, v.size, v.position, v.availability, v.price)
}

type Product struct{
  id int
  name string
  variants []Variant
}

func login(loginUrl, login string, password string) (requrl string) {
  fmt.Printf("Open: %s\n with: %s \n", loginUrl, password)

  doc, err := goquery.NewDocument(loginUrl)
  if err != nil {
    log.Fatal(err)
  }
  action, _ := doc.Find("#LoginBox form").Attr("action")
  // fmt.Printf("Action: %s\n", action)
  requrl, _ = doc.Find("body").Attr("requrl")
  // fmt.Printf("Url: %s\n", requrl)

  resp, err := http.PostForm(HOST + action, url.Values{
    "inpLoginUser": {login},
    "inpLoginPass": {password},
  })
  if err != nil {
    log.Fatalf("unable to fetch: %v", err)
  }
  defer resp.Body.Close()

  return
}

func loadProductUrl(productUrl string, productId string) []byte {
  productUrl = productUrl + "&monum=" + productId
  fmt.Printf("Open Product: %s\n", productUrl)

  resp, err := http.Get(productUrl)
  if err != nil {
    log.Fatalf("unable to fetch: %v", err)
  }
  defer resp.Body.Close()

  body, _ := ioutil.ReadAll(resp.Body)
  return body
}

func fileFor(productId string) string {
 return fmt.Sprintf("examples/product%s.html", productId)
}

func storeProductPage(productPage []byte, productId string) {
  path := fileFor(productId)
  ioutil.WriteFile(path, productPage, 0644)
}

func loadProductPage(productId string) []byte {
  path := fileFor(productId)
  file, _ := ioutil.ReadFile(path)
  return file
}

func inspectProduct(productPage []byte, productId string) Product {
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
  }

  reader := s.NewReader(string(productPage))
  doc, err := goquery.NewDocumentFromReader(reader)
  if err != nil {
    log.Fatal(err)
  }

  title := doc.Find(fmt.Sprintf("#styledesc%s b", productId)).Text()

  doc.Find(".tblTrArtRow").Each(func(i int, productSelection *goquery.Selection) {
    color    := func(value string) string {
      return string(value[0:3])
    }(productSelection.Find("td b").Text())

    priceStr := productSelection.Next().Find("b").Text()
    price, _ := strconv.ParseFloat(priceStr, 32)

    productSelection.Next().Find("input[type=text]").Each(func(i2 int, variantSelection *goquery.Selection) {
      size := func(value string, exists bool) string {
        splitAry := s.Split(value, "_")
        return sizeMapping[splitAry[len(splitAry)-1]]
      }(variantSelection.Attr("name"))

      availability := func(value string, exists bool) int {
          return availabilityMapping[value]
      }(variantSelection.Attr("class"))

      v := Variant {
        color: color,
        size: size,
        price: price,
        availability: availability,
      }

      fmt.Printf("V: %v\n", v.String())
    })
  })

  return Product {
    name: title,
  }
}

func main() {
  var productPage []byte;

  productId := os.Args[1]  // "X2-A", "EPIC", "F300"

  if len(os.Args) != 3 {
    productPage = loadProductPage(productId)
  } else {
    session := login(START_URL, LOGIN, os.Args[2])
    fmt.Printf("Session: %s\n", session)
    productPage = loadProductUrl(HOST + session, productId)
    storeProductPage(productPage, productId)
  }

  product := inspectProduct(productPage, productId)

  fmt.Printf("Product: %s\n", product)
}
