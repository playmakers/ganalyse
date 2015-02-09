package ganalyse

import (
  s "strings"
  "fmt"
  "net/http"
  "io/ioutil"
  "os"
  "encoding/csv"
  "path/filepath"
)

type Entry struct {
  Name string
  id, variantId int
  Shops map[string]string
}

// ------------------------------------------

func pathFor(shop string) string {
 shop = s.ToLower(s.Replace(shop, " ", "", -1))
 return fmt.Sprintf("examples/%s", shop)
}

func fileFor(productId string) string {
 productId = s.ToLower(s.Replace(productId, " ", "", -1))
 return fmt.Sprintf("%s.html", productId)
}

func loadUrl(url string) []byte {
  resp, err := http.Get(url)
  defer resp.Body.Close()

  // TODO handle 404 / and other errors
  if err != nil {
   // handle error

  }

  body, _ := ioutil.ReadAll(resp.Body)
  return body
}

func loadFile(path string) []byte {
  file, _ := ioutil.ReadFile(path)
  return file
}

func storeFile(fileName string, content []byte) {
  fmt.Printf(" --> store to: %s \n", fileName)
  err := ioutil.WriteFile(fileName, content, 0644)
  if os.IsNotExist(err) {
    path := filepath.Dir(fileName)
    os.MkdirAll(path, 0755)
    storeFile(fileName, content)
  }
}

func readCsv(fileName string) (records [][]string) {
  csvFile, _ := os.Open(fileName)
  defer csvFile.Close()

  reader := csv.NewReader(csvFile)
  // reader.FieldsPerRecord = -1

  records, _ = reader.ReadAll()
  return
}

// ------------------------------------------
func StoreUrl(shop string, productId string, url string) {
  filename := filepath.Join(pathFor(shop), fileFor(productId))
  if _, err := os.Stat(filename); os.IsNotExist(err) {
    fmt.Printf("Processing: %s, %s\n", filename, url)
    storeFile(filename, loadUrl(url))
  }
}

func ParseCsv(fileName string) (entries []Entry) {
  records := readCsv(fileName)
  for _, each := range records[1:len(records)] {
    entry := Entry{}
    entry.Name = each[1]
    entry.Shops = make(map[string]string)

    for i, url := range each[4:len(each)] {
      if url != "-" && url != "???" && !s.Contains(url, "keine") {
        entry.Shops[records[0][i+4]] = url
      }
    }
    entries = append(entries, entry)
  }

  return
}

// ------------------------------------------
