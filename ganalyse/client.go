// curl is a simple cURL replacement.
package main

import (
  "fmt"
  "log"
  "os"
  "io"
  "io/ioutil"
  "strings"
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
  price int
  availability int
}

type Product struct{
  id int
  name string
  variants []Variant
}

func login(loginUrl, login string, password string) string {
  fmt.Printf("Open: %s\n with: %s \n", loginUrl, password)

  doc, err := goquery.NewDocument(loginUrl)
  if err != nil {
    log.Fatal(err)
  }
  action, _ := doc.Find("#LoginBox form").Attr("action")
  fmt.Printf("Action: %s\n", action)
  requrl, _ := doc.Find("body").Attr("requrl")
  fmt.Printf("Url: %s\n", requrl)

  resp, err := http.PostForm(HOST + action, url.Values{
    "inpLoginUser": {login},
    "inpLoginPass": {password},
  })
  if err != nil {
    log.Fatalf("unable to fetch: %v", err)
  }
  defer resp.Body.Close()

  return requrl
}

func loadProductUrl(productUrl string, productId string) io.Reader {
  productUrl = productUrl + "&monum=" + productId
  fmt.Printf("Open Product: %s\n", productUrl)

  resp, err := http.Get(productUrl)
  if err != nil {
    log.Fatalf("unable to fetch: %v", err)
  }
  // defer resp.Body.Close()

  // body, _ := ioutil.ReadAll(resp.Body)
  // return body

  return resp.Body
}

func storeProductPage(productPage io.Reader, productId string) {
  path := fmt.Sprintf("product%s.html", productId)
  body, _ := ioutil.ReadAll(productPage)
  ioutil.WriteFile(path, body, 0644)
}

func loadProductPage(productId string) io.Reader {
  path := fmt.Sprintf("product%s.html", productId)
  file, _ := ioutil.ReadFile(path)
  return strings.NewReader(string(file))
}

func inspectProduct(productPage io.Reader, productId string) Product {
  doc, err := goquery.NewDocumentFromReader(productPage)
  if err != nil {
    log.Fatal(err)
  }

  title := doc.Find(fmt.Sprintf("#styledesc%s b", productId)).Text()

  return Product {
    name: title,
  }
}

func main() {
  if len(os.Args) != 2 {
    log.Fatalf("usage: %v pwd", os.Args[0])
  }
  productId := "EPIC"

  // session := login(START_URL, LOGIN, os.Args[1])
  // fmt.Printf("Session: %s\n", session)
  // productPage := loadProductUrl(HOST + session, productId)
  // storeProductPage(productPage, productId)

  productPage := loadProductPage(productId)

  product     := inspectProduct(productPage, productId)

  fmt.Printf("Product: %s\n", product)
}
