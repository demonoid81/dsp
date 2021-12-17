package utils

import (
	"github.com/mileusna/useragent"
	"net"
	"net/http"
	"strings"
	"github.com/oschwald/geoip2-golang"
)

func GetPlatform(useragent string) string {
	ua := ua.Parse(useragent)
	return ua.OS
}

func GetBrowser(useragent string) string {
	ua := ua.Parse(useragent)
	return ua.Name
}

func GetIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}

	s := strings.Split(IPAddress, ":")
	ip := s[0]
	return ip
}

func GetUA(r *http.Request) string {
	ua := r.Header.Get("User-Agent")
	return ua
}

func GetCountry(ip string) string {

	db, err := geoip2.Open("/var/www/GoApplications/src/GeoLite2-City.mmdb")
	if err != nil {
		return "RU"
	}
	defer db.Close()

	ip_addr := net.ParseIP(ip)
	record, err := db.City(ip_addr)
	if err != nil {
		return "RU"
	}

	return record.Country.IsoCode

}

func GetOS(useragent string) string {
	ua := ua.Parse(useragent)
	return ua.OS
}