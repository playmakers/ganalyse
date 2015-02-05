package main

import (
  "fmt"
  "sort"
  "log"
  "strings"
  "github.com/PuerkitoBio/goquery"
  // "database/sql"
  // _ "github.com/go-sql-driver/mysql"
)

// import "net/http"
// resp, err := http.Get("http://example.com/")
// if err != nil {
//   // handle error
// }
// defer resp.Body.Close()
// body, err := ioutil.ReadAll(resp.Body)


func download(targetUrl) Siegel {
  doc, err := goquery.NewDocument(fmt.Sprintf(tagetUrl, id))
  if err != nil {
    log.Fatal(err)
}


  siegel := Siegel{
    id: id,
    title: doc.Find(".sealdetaildesc h1").Text(),
    desc: doc.Find(".sealdetaildesc span").Text(),
    owner: doc.Find("#siegelDetailsBetreiber").Text(),
    category: doc.Find("#siegelDetailsKategorie").Text(),
    url: doc.Find("#siegelDetailsUrl").Text(),
    more: doc.Find("#siegelDetailsFreitext").Text(),
  }
  image, ok := doc.Find(".sealdetailmainpic img").Attr("src")

  if ok {
    siegel.image = image
  } else {
    siegel.image = "?"
  }

  siegel.regions = doc.Find("#siegelDetailsRegionen li").Map(func(i int, s *goquery.Selection) string {
    return s.Text()
  })

  siegel.branchen = doc.Find("#siegelDetailsBranchen li").Map(func(i int, s *goquery.Selection) string {
    return s.Text()
  })

  siegel.shops = doc.Find("#siegelDetailsUnternehmen li").Map(func(i int, s *goquery.Selection) string {
    return s.Text()
  })

  fmt.Printf("### Siegel %d: %s, %s, %s\n\n%s\n\n  %s\n  %s\n\n%s\n%s\n%s\n\n",
    siegel.id,
    siegel.title,
    siegel.owner,
    siegel.category,
    siegel.desc,
    siegel.url,
    siegel.image,
    strings.Join(siegel.regions, ","),
    strings.Join(siegel.branchen, ","),
    strings.Join(siegel.shops, ","),
  )

  return siegel
}

func main() {
  // db, err := sql.Open("mysql", "root:password@/dbname")
  // if err != nil {
  //     panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
  // }
  // defer db.Close()

  for i := 0; i < 20; i++ {
    if !skip(i) {
      go scrape(i)
    }
  }

  // prevent main from exiting immediately
  var input string
  fmt.Scanln(&input)
}


// stmtIns, err := db.Prepare("INSERT INTO siegels VALUES( ?, ? )") // ? = placeholder
// if err != nil {
//     panic(err.Error()) // proper error handling instead of panic in your app
// }
// defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
