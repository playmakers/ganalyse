package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/playmakers/ganalyse/lib/ganalyse"
	"github.com/playmakers/ganalyse/lib/vendors"
	"net/http"
	"os"
	"strings"
)

var port = os.Getenv("PORT")

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/extract", func(req *http.Request, r render.Render) {
		params := req.URL.Query()
		products := map[string]*vendors.Product{}
		urls := strings.Split(params.Get("url"), "\n")

		sem := make(chan bool, len(urls))

		for _, url := range urls {
			go func(url string) {
				products[url] = ganalyse.InspectUrl(url)
				sem <- true
			}(url)
		}

		for _, _ = range urls {
			<-sem
		}

		r.JSON(200, products)
	})

	if port == "" {
		port = "3000"
	}

	m.RunOnAddr(":" + port)
}
