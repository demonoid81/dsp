package dsp

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/demonoid81/dsp/auction/dsp/labyrinthads"
	"github.com/demonoid81/dsp/config"
	"github.com/demonoid81/dsp/events/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

var Affiliates = map[string]interface{}{
	"labyrinthads":  labyrinthads.Get,
}

type single struct {
	mu     sync.Mutex
	values map[string]string
}

var counters = single{
	values: make(map[string]string),
}

func (s *single) Get(key string) int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	var count int64
	nowDate := time.Now()
	timestamp := nowDate.Unix()
	if s.values[key] != "" {
		vals := strings.Split(s.values[key], "_")
		if vals[0] == strconv.FormatInt(timestamp, 10) {
			number, _ := strconv.ParseInt(vals[1], 10, 32)
			count = number
		} else {
			s.values[key] = strconv.FormatInt(timestamp, 10) + "_0"
			count = 0
		}
	} else {
		s.values[key] = strconv.FormatInt(timestamp, 10) + "_0"
		count = 0
	}
	return count
}

func (s *single) Set(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	vals := strings.Split(s.values[key], "_")
	number, _ := strconv.ParseInt(vals[1], 10, 32)
	count := number + 1
	nowDate := time.Now()
	timestamp := nowDate.Unix()
	s.values[key] = strconv.FormatInt(timestamp, 10) + "_" + strconv.FormatInt(count, 10)
}

type Creatives []struct {
	Body        string  `json:"body"`
	Cpc         float64 `json:"cpc"`
	CpcOriginal float64 `json:"cpc_original"`
	Icon        string  `json:"icon"`
	ID          string  `json:"id"`
	Image       string  `json:"image"`
	Link        string  `json:"link"`
	Status      int     `json:"status"`
	Title       string  `json:"title"`
	DspId       string  `json:"dsp_id"`
	DspName     string  `json:"dsp_name"`
	SspId       string  `json:"ssp_id"`
	SspName     string  `json:"ssp_name"`
	Timestamp   int     `json:"timestamp"`
}

func Event(ctx context.Context, data map[string]interface{}, cfg map[string]interface{}, waitGroup *sync.WaitGroup, mongoClient *mongo.Client) (map[string]interface{}, string) {

	c := make(chan map[string]interface{})
	var creatives []map[string]interface{}
	result := Creatives{}

	for _, dsp := range cfg["dsp"].([]map[string]interface{}) {
		var dataDsp = map[string]interface{}{
			"dsp_id":   dsp["dsp_id"],
			"dsp_name": config.Config["DSP"].(map[string]interface{})[dsp["dsp_id"].(string)].(map[string]interface{})["dsp_name"],
			"profit":   dsp["profit"],
			"endpoint": config.Config["DSP"].(map[string]interface{})[dsp["dsp_id"].(string)].(map[string]interface{})["endpoint"],
			"ssp_id":   cfg["ssp_id"],
			"ssp_name": cfg["ssp_name"],
		}

		if (!utils.ContainsInArray(dsp["source_id_blacklist"], data["sid"].(string)) &&
			!utils.ContainsInArray(dsp["country_blacklist"], data["country"].(string)) &&
			(len(reflect.ValueOf(dsp["country_whitelist"]).Interface().([]string)) == 0 || utils.ContainsInArray(dsp["country_whitelist"], data["country"].(string)))) &&
			int(counters.Get(dsp["dsp_id"].(string))) <= (config.Config["DSP"].(map[string]interface{})[dsp["dsp_id"].(string)].(map[string]interface{})["qps"].(int)/2) {

			counters.Set(dsp["dsp_id"].(string))
			data["sid"] = data["sid"].(string)
			data["id"] = dsp["dsp_id"].(string) + cfg["ssp_id"].(string) + data["sid"].(string)

			go Affiliates[dataDsp["dsp_name"].(string)].
				(func(context.Context, map[string]interface{}, chan map[string]interface{}, map[string]interface{}, *sync.WaitGroup,  *mongo.Client))(ctx, data, c, dataDsp, waitGroup, mongoClient)
			creatives = append(creatives, <-c)

		}
	}

	close(c)

	res, err := json.Marshal(creatives)
	if err != nil {
		return map[string]interface{}{
			"status": 204,
		}, ""
	}

	json.Unmarshal(res, &result)

	sort.Slice(result, func(i, j int) bool {
		return result[i].Cpc > result[j].Cpc
	})

	var creative []map[string]interface{}
	inrec, _ := json.Marshal(result)
	json.Unmarshal(inrec, &creative)

	returnCreative := map[string]interface{}{}

	if len(creative) > 0 {
		returnCreative = creative[0]
	}

	var headerCreatives []map[string]interface{}

	for index, crKafka := range creative {
		cr_ := map[string]interface{}{
			"status": crKafka["status"],
			"d_id":   crKafka["dsp_id"],
		}
		if fmt.Sprint(crKafka["status"]) == "200" {
			if index == 0 {
				cr_["w"] = 1
				cr_["cpc"] = crKafka["cpc_original"]
			}
		}
		headerCreatives = append(headerCreatives, cr_)
	}

	/*header_data := map[string]interface{}{
	      "ssp_id": ssp_key,
	      "sub_id": data["sid"],
	      "country": data["country"],
	      "data": header_creatives,
	  }

	  json_header, _ := json.Marshal(header_data)
	*/
	jsonHeader, _ := json.Marshal(headerCreatives)
	base64Header := base64.StdEncoding.EncodeToString(jsonHeader)

	return returnCreative, base64Header

}

