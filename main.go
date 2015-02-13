package main

import (
  "./ganalyse"
  "sync"
  "fmt"
  // "os/exec"
)

func parse(shop string, data []byte) {
  switch shop {
    case "1A": {
      // product := ganalyse.Inspect1A(data)
      // fmt.Printf(" --> Product: %v\n", product.String())
    }
    case "Boss Hogg": {
      // product := ganalyse.InspectBossHogg(data)
      // fmt.Printf(" --> Product: %v\n", product.String())
    }
    case "DocA": {
      product := ganalyse.InspectBossDocA(data)
      fmt.Printf(" --> Product: %v\n", product.String())
    }
    case "First Down": {
      // product := ganalyse.Inspect1A(data)
      // fmt.Printf(" --> Product: %v\n", product.String())
    }
    case "Forelle": {
      ganalyse.InspectForelle(data)
      // fmt.Printf(" --> Product: %v\n", product.String())
    }
    case "Futspo": {
      ganalyse.InspectFutspo(data)
      // fmt.Printf(" --> Product: %v\n", product.String())
    }
    case "Meyer": {
      ganalyse.InspectMeyer(data)
      // fmt.Printf(" --> Product: %v\n", product.String())
    }
    case "Playmakers": {
      ganalyse.InspectPotsdam(data)
      // fmt.Printf(" --> Product: %v\n", product.String())
    }
    case "Potsdam": {
      ganalyse.InspectPotsdam(data)
      // fmt.Printf(" --> Product: %v\n", product.String())
    }
    case "Sports and Cheer": {
      // product := ganalyse.Inspect1A(data)
      // fmt.Printf(" --> Product: %v\n", product.String())
    }
    case "US Sportshop": {
      // product := ganalyse.Inspect1A(data)
      // fmt.Printf(" --> Product: %v\n", product.String())
    }
    default: {
      fmt.Printf(" -> %s\n", shop)
    }
  }
}

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
    fmt.Printf("%s\n", entry.Name)
    for shop, url := range entry.Shops {
      //*** download(targetUrl)
      wg.Add(1)
      func(shop string, name string, url string){
        defer wg.Done()
        filename := ganalyse.StoreUrl(shop, entry.Name, url)
        data := ganalyse.LoadFile(filename)
        parse(shop, data)
      }(shop, entry.Name, url)
    }
  }

  wg.Wait()

  //*** map to pageobject, analyse content


  //*** send to ga with custom dimensions/value


  //*** update google doc with values & status
}
