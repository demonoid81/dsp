package useragent

import (
	"github.com/mileusna/useragent"
)

func GetPlatform(useragent string) string {
	ua := ua.Parse(useragent)
	return ua.OS
}

func GetBrowser(useragent string) string {
	ua := ua.Parse(useragent)
	return ua.Name
}
