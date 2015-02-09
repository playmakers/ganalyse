package main

import (
  "./ganalyse"
  "sync"
  "fmt"
  // "os/exec"
)

func main() {
  // if len(os.Args) != 1 {
  //   log.Fatalf("usage: %v $URL", os.Args[0])
  // }
  // if _, err := http.Get(os.Stdout, os.Args[1]); err != nil {
  //   log.Fatalf("unable to fetch %q: %v", os.Args[1], err)
  // }

  // TODO execute speadsheet downloader

  var wg sync.WaitGroup

  //*** read file/google doc
  entries := ganalyse.ParseCsv("examples/data.csv")

  //*** get data from given urls
  for _, entry := range entries {
    for shop, url := range entry.Shops {
      //*** download(targetUrl)
      wg.Add(1)
      func(shop string, name string, url string){
        defer wg.Done()
        fmt.Printf("%s %s %s\n", shop, name, url)
        ganalyse.StoreUrl(shop, entry.Name, url)
      }(shop, entry.Name, url)
    }
  }

  wg.Wait()

  //*** map to pageobject, analyse content


  //*** send to ga with custom dimensions/value


  //*** update google doc with values & status
}
