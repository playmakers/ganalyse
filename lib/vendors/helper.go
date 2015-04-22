package vendors

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
	"io"
	"regexp"
	s "strings"
)

type sizeWithPrice struct {
	size  string
	price float64
}

type mapping func(string) string

const (
	DEFAULT_SIZE         = "L"
	DEFAULT_COLOR        = "Schwarz"
	DEFAULT_AVAILABILITY = 10
)

func Parse(data []byte, charset string) (doc *goquery.Document) {
	var reader io.Reader
	reader = s.NewReader(string(data))
	if charset != "utf-8" {
		reader, _ = iconv.NewReader(reader, charset, "utf-8")
	}
	doc, _ = goquery.NewDocumentFromReader(reader)
	return
}

func findOption(haystack *goquery.Selection, needle string) *goquery.Selection {
	return haystack.FilterFunction(func(i int, selection *goquery.Selection) bool {
		return selection.Parent().Prev().Text() == needle
	}).Find("option")
}

func dropFirst(selection *goquery.Selection) *goquery.Selection {
	if selection.Size() > 0 {
		return selection.Slice(1, -1)
	} else {
		return selection.Remove()
	}
}

func getValues(selection *goquery.Selection, defaultValue string, mapping mapping) (values []string) {
	selection.Each(func(i int, valueSelection *goquery.Selection) {
		value := func(value string, exists bool) string {
			return mapping(s.TrimSpace(value))
		}(valueSelection.Attr("value"))

		values = append(values, value)
	})
	if len(values) < 1 {
		values = append(values, defaultValue)
	}
	return
}

func getSizes(selection *goquery.Selection, regMatcher *regexp.Regexp) (sizes []sizeWithPrice) {
	selection.Each(func(i int, sizeSelection *goquery.Selection) {
		sizeString := sizeSelection.Text()
		r := regMatcher.FindAllStringSubmatch(sizeString, -1)
		if len(r) > 0 {
			sizes = append(sizes, sizeWithPrice{
				size:  r[0][1],
				price: NormPrice(r[0][3]),
			})
		}
	})
	if len(sizes) < 1 {
		sizes = append(sizes, sizeWithPrice{
			size:  DEFAULT_SIZE,
			price: 0,
		})
	}
	return
}

func getColors(selection *goquery.Selection) (colors []string) {
	selection.Each(func(i int, colorSelection *goquery.Selection) {
		colors = append(colors, func(value string) string {
			return s.TrimSpace(value)
		}(colorSelection.Text()))
	})
	if len(colors) < 1 {
		colors = append(colors, DEFAULT_COLOR)
	}
	return
}
