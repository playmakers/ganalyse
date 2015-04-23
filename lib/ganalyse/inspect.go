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

func InspectUrl(raw_url string) *vendors.Product {
	re := regexp.MustCompile("www\\.")
	url, err := url.Parse(raw_url)
	if err != nil {
		return nil
	}
	if url.Host != "" {
		domain := re.ReplaceAllString(url.Host, "")
		data, _ := loadUrl(raw_url)
		return parse(domain, data)
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
	return parse(filepath.Dir(path), data)
}

func parse(id string, data []byte) *vendors.Product {
	switch id {
	case "1a-football.de", "1a":
		{
			return vendors.Inspect1A(data)
		}
	case "tms-footballshop-berlin.de", "Boss Hogg":
		{
			return vendors.InspectBossHogg(data)
		}
	case "doca-sports.com", "DocA":
		{
			return vendors.InspectDocA(data)
		}
	case "mallux.de", "First Down":
		{
			return vendors.InspectFirstDown(data)
		}
	case "forelle.com", "Forelle":
		{
			return vendors.InspectForelle(data)
		}
	case "futspo.de", "Futspo":
		{
			return vendors.InspectFutspo(data)
		}
	case "dfshop.com", "Meyer":
		{
			return vendors.InspectMeyer(data)
		}
	// case "Playmakers": {
	//   return vendors.InspectPlaymakers(data)
	// }
	case "american-footballshop.de", "Potsdam":
		{
			return vendors.InspectPotsdam(data)
		}
	case "sportsandcheer.de", "Sports and Cheer":
		{
			return vendors.InspectSportsAndCheer(data)
		}
	case "us-sportshop.de", "US Sportshop":
		{
			return vendors.InspectUsSportshop(data)
		}
	default:
		{
			return nil
		}
	}
}
