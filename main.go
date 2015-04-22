package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
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
	m.Use(render.Renderer())

	m.Get("/extract", func(req *http.Request, r render.Render) {
		params := req.URL.Query()
		vendor := params.Get("vendor")
		url := params.Get("url")
		data, _ := ganalyse.LoadUrl(url)
		product := parse(vendor, data)
		if product != nil {
			r.JSON(200, product)
		} else {
			r.JSON(500, map[string]interface{}{"message": "error", "vendor": vendor, "url": url})
			// return "Couldn't parse " + vendor + " " + url
		}
	})

	if port == "" {
		port = "3000"
	}

	m.RunOnAddr(":" + port)
}
