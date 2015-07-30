package ganalyse

import (
	"errors"
	"github.com/playmakers/ganalyse/lib/vendors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"regexp"
)

func InspectUrl(rawUrl string) *vendors.Product {
	re := regexp.MustCompile("www\\.")
	url, err := url.Parse(rawUrl)
	if err != nil {
		return nil
	}
	if url.Host != "" {
		domain := re.ReplaceAllString(url.Host, "")
		data, _ := loadUrl(rawUrl)
		return parse(domain, data, rawUrl)
	} else {
		return nil
	}
}

func loadUrl(url string) ([]byte, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		log.Fatal("Couldn't load URL")
	}
	if resp.StatusCode == 404 {
		return nil, errors.New("Page not available")
	}

	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}

func InspectFile(path string) *vendors.Product {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Couldn't load URL")
		return nil
	}
	return parse(filepath.Dir(path), data, path)
}

func parse(id string, data []byte, origin string) *vendors.Product {
	switch id {
	case "1a-football.de", "1a":
		{
			return vendors.Inspect1A(data, origin)
		}
	case "tms-footballshop-berlin.de", "Boss Hogg":
		{
			return vendors.InspectBossHogg(data, origin)
		}
	case "doca-sports.com", "DocA":
		{
			return vendors.InspectDocA(data, origin)
		}
	case "mallux.de", "First Down":
		{
			return vendors.InspectFirstDown(data, origin)
		}
	case "forelle.com", "Forelle":
		{
			return vendors.InspectForelle(data, origin)
		}
	case "futspo.de", "Futspo":
		{
			return vendors.InspectFutspo(data, origin)
		}
	case "dfshop.com", "Meyer":
		{
			return vendors.InspectMeyer(data, origin)
		}
	// case "Playmakers": {
	//   return vendors.InspectPlaymakers(data)
	// }
	case "american-footballshop.de", "Potsdam":
		{
			return vendors.InspectPotsdam(data, origin)
		}
	case "sportsandcheer.de", "Sports and Cheer":
		{
			return vendors.InspectSportsAndCheer(data, origin)
		}
	case "us-sportshop.de", "US Sportshop":
		{
			return vendors.InspectUsSportshop(data, origin)
		}
	default:
		{
			return nil
		}
	}
}
