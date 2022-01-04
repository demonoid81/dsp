package timestamp

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func Compatible(timestamp string, freshness string) (bool) {

	iTimestamp, err := strconv.Atoi(timestamp)

	if err != nil {
		return false
	}

	aFreshness := map[string]interface{}{}

	json.Unmarshal([]byte(freshness), &aFreshness)

	fType := aFreshness["type"].(string)
	fInterval1 := int(aFreshness["interval1"].(float64))
	fInterval2 := int(aFreshness["interval2"].(float64))

	var sec int

	if fType == "m" {
		sec = 60
	} else if fType == "h" {
		sec = 3600
	} else if fType == "d" {
		sec = 86400
	}

	now := time.Now()
	currentTimestamp := int(now.Unix())

	interval1 := currentTimestamp - (fInterval1 * sec)
	interval2 := currentTimestamp - (fInterval2 * sec)

	if iTimestamp >= interval2 && iTimestamp <= interval1 {
		fmt.Println("TZ Compatible: true" )
		return true
	} else {
		fmt.Println("TZ Compatible: false" )
		return false
	}

}

func Freshness(timestamp string) string {

	iTimestamp, err := strconv.Atoi(timestamp)

	if err != nil {
		return "error"
	}

	now := time.Now()
	currentTimestamp := int(now.Unix())

	seconds := currentTimestamp - iTimestamp

	if seconds >= 0 && seconds <= (60*5) {
		return "0x5"
	}

	if seconds > (60*5) && seconds <= (60*10) {
		return "5x10"
	}

	if seconds > (60*10) && seconds <= (60*20) {
		return "10x20"
	}

	if seconds > (60*20) && seconds <= (60*30) {
		return "20x30"
	}

	if seconds > (60*30) && seconds <= (60*60) {
		return "30x60"
	}

	if seconds > (3600) && seconds <= (3*3600) {
		return "1y3"
	}

	if seconds > (3*3600) && seconds <= (12*3600) {
		return "3y12"
	}

	if seconds > (12*3600) && seconds <= (24*3600) {
		return "12y24"
	}

	if seconds > (86400) && seconds <= (3*86400) {
		return "1z3"
	}

	if seconds > (3*86400) && seconds <= (7*86400) {
		return "3z7"
	}

	if seconds > (7 * 86400) {
		return "7d"
	}

	return "error"
}
