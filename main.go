package main

import (
	"github.com/go-martini/martini"
	"net/http"
	// "log"
	"github.com/playmakers/ganalyse/lib/ganalyse"
	"github.com/playmakers/ganalyse/lib/vendors"
	"os"
)

func parse(shop string, data []byte) *ganalyse.Product {
	switch shop {
	case "1A":
		{
			return vendors.Inspect1A(data)
		}
	case "Boss Hogg":
		{
			return vendors.InspectBossHogg(data)
		}
	case "DocA":
		{
			return vendors.InspectDocA(data)
		}
	case "First Down":
		{
			return vendors.InspectFirstDown(data)
		}
	case "Forelle":
		{
			return vendors.InspectForelle(data)
		}
	case "Futspo":
		{
			return vendors.InspectFutspo(data)
		}
	case "Meyer":
		{
			return vendors.InspectMeyer(data)
		}
	// case "Playmakers": {
	//   return vendors.InspectPotsdam(data)
	// }
	case "Potsdam":
		{
			return vendors.InspectPotsdam(data)
		}
	case "Sports and Cheer":
		{
			return vendors.InspectSportsAndCheer(data)
		}
	case "US Sportshop":
		{
			return vendors.InspectUsSportshop(data)
		}
	default:
		{
			return nil
		}
	}
}

var port = os.Getenv("PORT")

func main() {
	m := martini.Classic()

	m.Get("/extract/:vendor", func(params martini.Params, req *http.Request) string {
		url := req.URL.Query().Get("url")
		data, _ := ganalyse.LoadUrl(url)
		product := parse(params["vendor"], data)
		if product != nil {
			return product.String()
		} else {
			return "Couldn't parse" + params["vendor"] + " " + url
		}
	})

	if port == "" {
		port = "3000"
	}

	m.RunOnAddr(":" + port)
}
