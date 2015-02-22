package main

import (
  "./lib/ganalyse"
  "./lib/tracker"
  "./lib/vendors"
  "sync"
  "fmt"
  // "os/exec"
)

func parse(shop string, data []byte) *ganalyse.Product {
  switch shop {
    case "1A": {
      return vendors.Inspect1A(data)
    }
    case "Boss Hogg": {
      return vendors.InspectBossHogg(data)
    }
    case "DocA": {
      return vendors.InspectDocA(data)
    }
    case "First Down": {
      return vendors.InspectFirstDown(data)
    }
    case "Forelle": {
      return vendors.InspectForelle(data)
    }
    case "Futspo": {
      return vendors.InspectFutspo(data)
    }
    case "Meyer": {
      return vendors.InspectMeyer(data)
    }
    // case "Playmakers": {
    //   return vendors.InspectPotsdam(data)
    // }
    case "Potsdam": {
      return vendors.InspectPotsdam(data)
    }
    case "Sports and Cheer": {
      return vendors.InspectSportsAndCheer(data)
    }
    case "US Sportshop": {
      return vendors.InspectUsSportshop(data)
    }
    default: {
      return nil
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
    fmt.Printf("\n\n%s\n", entry.Name)
    for shop, url := range entry.Shops {
      //*** download(targetUrl)
      wg.Add(1)
      func(shop string, name string, url string){
        defer wg.Done()
        filename := ganalyse.StoreUrl(shop, entry.Name, url)
        data := ganalyse.LoadFile(filename)
        product := parse(shop, data)
        if product != nil {
          fmt.Printf(" --> %s: %v\n", shop, product.String())
        }
      }(shop, entry.Name, url)
    }
  }

  wg.Wait()

  //*** map to pageobject, analyse content


  //*** send to ga with custom dimensions/value


  //*** update google doc with values & status
}
