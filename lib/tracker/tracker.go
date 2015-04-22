package tracker

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"github.com/playmakers/ganalyse/lib/ganalyse"
	"net/http"
	"net/url"
	"strconv"
	s "strings"
)

const (
	NUMBER = "UA-45979504-2"
	URL    = "http://www.google-analytics.com/collect"
)

func send1(url string, values url.Values) {
	fmt.Println(URL + "?" + values.Encode())

	v, err := http.PostForm(url, values)
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}

func send2(url string, values url.Values) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, s.NewReader(values.Encode()))
	req.Header.Add("User-Agent", "ganalyse")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
}

func Track(shop string, trackingId int64, productId int, variantId int, product *vendors.Product, variant *vendors.Variant) {
	vals := make(url.Values, 0)
	vals.Add("v", "1")
	vals.Add("ds", "ganalyse")
	vals.Add("tid", NUMBER)
	vals.Add("cid", uuid.NewRandom().String())
	vals.Add("t", "pageview")
	vals.Add("dp", "/")
	vals.Add("uid", shop) // opt.
	vals.Add("ci", fmt.Sprintf("%d", trackingId))
	vals.Add("cn", fmt.Sprintf("%d", trackingId))
	vals.Add("cm", "auto")
	vals.Add("cs", "direct")
	vals.Add("ck", shop)

	vals.Add("cd1", shop)
	// vals.Add("cd2", productType)
	// vals.Add("cd3", product.Name)
	// vals.Add("cd4", variant.Size)
	// vals.Add("cd5", variant.Color)
	// vals.Add("cd6", variant.Pos)
	vals.Add("cd7", strconv.Itoa(productId))
	vals.Add("cd8", strconv.Itoa(variantId))
	// vals.Add("cd9", product.Vendor)
	vals.Add("cm1", fmt.Sprintf("%.2f", variant.Price))
	vals.Add("cm2", strconv.Itoa(variant.Availability))

	// send2(URL, vals)
}
