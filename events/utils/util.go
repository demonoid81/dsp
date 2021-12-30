package utils

import (
	"github.com/mileusna/useragent"
	"github.com/oschwald/geoip2-golang"
	"net"
	"net/http"
	"reflect"
	"strings"
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


// todo расширить до версии
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

func ContainsInArray(t interface {}, str string) bool {

	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)

		for i := 0; i < s.Len(); i++ {
			if s.Index(i).Interface().(string) == str {
				return true
			}
		}
	}

	return false
}


func Find(slice interface{}, f func(value interface{}) bool) int {
	s := reflect.ValueOf(slice)
	if s.Kind() == reflect.Slice {
		for index := 0; index < s.Len(); index++ {
			if f(s.Index(index).Interface()) {
				return index
			}
		}
	}
	return -1
}
