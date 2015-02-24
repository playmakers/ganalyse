package tracker

import (
  "fmt"
  "net/http"
  "net/url"
  "strings"
  "../ganalyse"
)

const (
 NUMBER = "UA-45979504-2"
 URL  = "http://www.google-analytics.com/collect"
)

func send1(url string, values url.Values) {
 fmt.Println(URL + "?" + values.Encode())

  v, err := http.PostForm(url, values);
  if err != nil {
    panic(err)
  }
  fmt.Println(v)
}

func send2(url string, values url.Values) {
  client := &http.Client{}
  req, _ := http.NewRequest("POST", url, strings.NewReader(values.Encode()))
  req.Header.Add("User-Agent", "myClient")
  resp, _ := client.Do(req)
  defer resp.Body.Close()
}

func Track(shop string, productType string, product *ganalyse.Product, variant *ganalyse.Variant) {
  vals := make(url.Values, 0)
  vals.Add("v", "1")
  vals.Add("tid", NUMBER)
  vals.Add("cid", "31fe906c-2aac-4821-a1ee-d0a9a09c2e0b")
  vals.Add("t", "pageview")
  vals.Add("dp", "/")
  vals.Add("cd1", shop)
  vals.Add("cd2", productType)
  vals.Add("cd3", product.Name)
  vals.Add("cd4", variant.Size)
  vals.Add("cd5", variant.Color)
  // vals.Add("cd6", variant.Pos)
  // vals.Add("cd7", product.Vendor)
  // vals.Add("cd8", product.Id)
  // vals.Add("cd9", variant.Id)
  vals.Add("cm1", fmt.Sprintf("%.2f", variant.Price))
  vals.Add("cm2", fmt.Sprintf("%d", variant.Availability))

  send2(URL, vals)
}
