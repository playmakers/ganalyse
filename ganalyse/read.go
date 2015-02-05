package ganalyse

import (
  "encoding/csv"
  "fmt"
  "os"
)

type Entry struct {
  productName, shop, url string
  productId, variantId int
}


func open(fileName string) {
  csvfile, err := os.Open(fileName)

  if err != nil {
    fmt.Println(err)
    return
  }

  defer csvfile.Close()

  reader := csv.NewReader(csvfile)

  reader.Comma = ';'
  reader.FieldsPerRecord = -1 // see the Reader struct information below

  rawCSVdata, err := reader.ReadAll()

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  return rawCSVdata
}

func Read(fileName string) {
  rawCSVdata := open(fileName)

  // sanity check, display to standard output
  for _, each := range rawCSVdata {
    fmt.Printf("Name: %s Forelle: %s\n", each[1], each[4])
  }
}
