package tracker

import (
  "net/http"
  "net/url"
)

func Track() {

  vals := make(url.Values, 0)
  vals.Add("v", "1")
  vals.Add("tid", "UA-XXXXXX-1")
  vals.Add("cid", "555")
  vals.Add("t", "event")
  vals.Add("ec", "Part")
  vals.Add("ea", "GetPart")
  vals.Add("el", "dataType=JSON&partID=11000")
  vals.Add("ev", "200")

  if _, err := http.Get("http://www.google-analytics.com/collect?" + vals.Encode()); err != nil {
    panic(err)
  }

}
