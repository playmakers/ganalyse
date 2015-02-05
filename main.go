package main

import (
  "./ganalyse"
)

const (
  ENDPOINT string = 'http://www.dfshop.com/dfshop/wsDfShop.wsc/pDfStart.p?href=dqbg4p735uvb12kccosjqb6'
  LOGIN string = 'PMBE6R47'
  PASSWORD string = 'IFKW2BDS'
)

func main() {
  if len(os.Args) != 1 {
    log.Fatalf("usage: %v $URL", os.Args[0])
  }
  // if _, err := http.Get(os.Stdout, os.Args[1]); err != nil {
  //   log.Fatalf("unable to fetch %q: %v", os.Args[1], err)
  // }

resp, err := http.PostForm("http://example.com/form",
  url.Values{"key": {"Value"}, "id": {"123"}})

  // read file/google doc
  // ganalyse.Read("data.csv")


  // get data from given urls

  // download(targetUrl)

  // map to pageobject, analyse content


  // send to ga with custom dimensions/value


  // update google doc with values & status
}
