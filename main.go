package main

import (
  "./lib/ganalyse"
  "./lib/tracker"
  "./lib/vendors"
  // "sync"
  "fmt"
  "time"
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

  // var wg sync.WaitGroup

  // //*** read file/google doc
  entries := ganalyse.ParseCsv("examples/data.csv")

  trackingId := time.Now().Unix()

  //*** get data from given urls
  for _, entry := range entries {
    fmt.Printf("\n\n%s\n", entry.Name)
    for shop, url := range entry.Shops {

      //*** download(targetUrl)
      // wg.Add(1)
      func(productShop string, productId int, variantId int, productName string, productUrl string){
        // defer wg.Done()
        filename := ganalyse.StoreUrl(productShop, productName, productUrl)
        data := ganalyse.LoadFile(filename)
        product := parse(productShop, data)
        if product != nil {
          fmt.Printf(" --> %s: %v\n", productShop, product.String())
          tracker.Track(
            productShop,
            trackingId,
            productId,
            variantId,
            product,
            product.DefaultVariant(),
          )
        }
      }(shop, entry.Id, entry.VariantId, entry.Name, url)
    }
  }

  fmt.Println(trackingId)

  // wg.Wait()

  //*** map to pageobject, analyse content


  //*** send to ga with custom dimensions/value

  // tracker.Track(
  //   "Forelle",
  //   "Handschuh",
  //   ganalyse.Product {
  //     // Id = 1,
  //     // Vendor: "Test",
  //     Name: "Superbad 3.0",
  //   },
  //   ganalyse.Variant {
  //     // Id = 1,
  //     Size: "L",
  //     Color: "schwarz",
  //     Price: 12.13,
  //     Availability: 10,
  //   },
  // )

  //*** update google doc with values & status
}
