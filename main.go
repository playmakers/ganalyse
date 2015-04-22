package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/playmakers/ganalyse/lib/ganalyse"
	"net/http"
	"os"
)

var port = os.Getenv("PORT")

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/extract", func(req *http.Request, r render.Render) {
		params := req.URL.Query()
		url := params.Get("url")
		product := ganalyse.Inspect(url)
		if product != nil {
			r.JSON(200, product)
		} else {
			r.JSON(500, map[string]interface{}{"message": "error", "url": url})
		}
	})

	if port == "" {
		port = "3000"
	}

	m.RunOnAddr(":" + port)
}
