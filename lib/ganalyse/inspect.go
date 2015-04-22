package ganalyse

import (
	"github.com/playmakers/ganalyse/lib/vendors"
	"net/url"
	"regexp"
)

func Inspect(raw_url string) *vendors.Product {
	re := regexp.MustCompile("www\\.")
	data, _ := LoadUrl(raw_url)
	url, _ := url.Parse(raw_url)
	domain := re.ReplaceAllString(url.Host, "")
	return parse(domain, data)
}

func parse(domain string, data []byte) *vendors.Product {
	switch domain {
	case "1a-football.de": //"1a":
		{
			return vendors.Inspect1A(data)
		}
	case "tms-footballshop-berlin.de": //"Boss Hogg":
		{
			return vendors.InspectBossHogg(data)
		}
	case "doca-sports.com": //"DocA":
		{
			return vendors.InspectDocA(data)
		}
	case "mallux.de": //"First Down":
		{
			return vendors.InspectFirstDown(data)
		}
	case "forelle.com": //"Forelle":
		{
			return vendors.InspectForelle(data)
		}
	case "futspo.de": //"Futspo":
		{
			return vendors.InspectFutspo(data)
		}
	case "dfshop.com": //"Meyer":
		{
			return vendors.InspectMeyer(data)
		}
	// case "Playmakers": {
	//   return vendors.InspectPlaymakers(data)
	// }
	case "american-footballshop.de": //"Potsdam":
		{
			return vendors.InspectPotsdam(data)
		}
	case "sportsandcheer.de": //"Sports and Cheer":
		{
			return vendors.InspectSportsAndCheer(data)
		}
	case "us-sportshop.de": //"US Sportshop":
		{
			return vendors.InspectUsSportshop(data)
		}
	default:
		{
			return nil
		}
	}
}
