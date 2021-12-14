package main

import (
	"context"
	"fmt"
	"github.com/demonoid81/dsp/events/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"time"

	//"os"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"strings"

	"github.com/demonoid81/dsp/config"
	"github.com/demonoid81/dsp/events/encrypt"
	"github.com/demonoid81/dsp/events/inArray"
	"github.com/demonoid81/dsp/events/postgres"
	"github.com/demonoid81/dsp/events/redis"
	ts "github.com/demonoid81/dsp/events/timestamp"
	"github.com/demonoid81/dsp/events/useragent"
	"github.com/google/uuid"
)

type campany struct {
	UID float64 `json:"uid"`
	Cur string  `json:"cur"`
	Cpr float64 `json:"cpr"`
	Cid string  `json:"cid"`
	Atl string  `json:"atl"`
	Atx string  `json:"atx"`
	Aic string  `json:"aic"`
	Aig string  `json:"aig"`
	Ccr string  `json:"ccr"`
}

type LinkData struct {
	Key    uuid.UUID `json:"key" bson:"key"`
	Link   string    `json:"link" bson:"link"`
	Cpc    float64   `json:"cpc" bson:"cpc"`
	Uid    float64   `json:"uid" bson:"uid"`
	Cid    string    `json:"cid" bson:"cid"`
	Cou    string    `json:"cou" bson:"cou"`
	Bro    string    `json:"bro" bson:"bro"`
	Os     string    `json:"os" bson:"os"`
	Sid    string    `json:"sid" bson:"sid"`
	Date   string    `json:"date" bson:"date"`
	Fresh  string    `json:"fresh" bson:"fresh"`
	FeedId string    `json:"feed_id" bson:"feed_id"`
	Click  bool      `json:"-" bson:"key"`
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	var waitGroup sync.WaitGroup

	mongoClient, err := mongodb.NewClient()

	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/ssp", func(w http.ResponseWriter, req *http.Request) {

		rdb := redis.Client()

		params, err := url.PathUnescape(req.URL.RawQuery)
		if err != nil {
			w.WriteHeader(204)
			return
		}
		dataReq := url.Values{}
		parameters := strings.Split(params, "&")
		for _, parameter := range parameters {
			parts := strings.Split(parameter, "=")
			dataReq.Add(parts[0], parts[1])
		}

		ua := strings.ReplaceAll(dataReq["ua"][0], "+", " ")

		browser := useragent.GetBrowser(ua)
		platform := useragent.GetPlatform(ua)

		fmt.Println(dataReq)

		country := dataReq["country"][0]
		sourceId := dataReq["sid"][0]
		category := dataReq["category"][0]
		timestamp := dataReq["timestamp"][0]
		pushType := dataReq["push_type"][0]
		feedId := dataReq["feed_id"][0]

		var redisKey = browser + "-" + platform + "-" + country + "-" + category + "-" + pushType + "-" + feedId
		var campaignsJson string

		redisCampaigns := redis.Get(rdb, redisKey)

		if redisCampaigns != "error" {

			campaignsJson = redisCampaigns

		} else {

			postgresCampaigns := postgres.E(country, platform, browser, category, pushType)

			if postgresCampaigns != "error" {

				campaignsJson = postgresCampaigns

				set := redis.Set(rdb, redisKey, campaignsJson)

				if set == "error" {
					w.WriteHeader(204)
					return
				}

			} else {
				w.WriteHeader(204)
				return
			}

		}

		conn := rdb.Get()
		conn.Close()

		var campaignsMap = []map[string]interface{}{}
		json.Unmarshal([]byte(campaignsJson), &campaignsMap)

		var Campaigns []campany

		for _, cfgCompany := range campaignsMap {
			blacklist := []string{}
			whitelist := []string{}

			blacklistFeed := []string{}
			whitelistFeed := []string{}

			json.Unmarshal([]byte(cfgCompany["blacklist"].(string)), &blacklist)
			json.Unmarshal([]byte(cfgCompany["whitelist"].(string)), &whitelist)

			json.Unmarshal([]byte(cfgCompany["blacklist_feed"].(string)), &blacklistFeed)
			json.Unmarshal([]byte(cfgCompany["whitelist_feed"].(string)), &whitelistFeed)

			if ts.Compatible(timestamp, cfgCompany["freshness"].(string)) &&
				inArray.FindString(blacklist, sourceId) == false &&
				(len(whitelist) <= 0 || inArray.FindString(whitelist, sourceId) == true) &&
				inArray.FindString(blacklistFeed, feedId) == false &&
				(len(whitelistFeed) <= 0 || inArray.FindString(whitelistFeed, feedId) == true) {

				var _Campany = campany{
					UID: cfgCompany["user_id"].(float64),
					Cur: cfgCompany["company_url"].(string),
					Cpr: cfgCompany["company_price"].(float64),
					Cid: strconv.FormatFloat(cfgCompany["company_id"].(float64), 'f', -1, 64),
					Atl: cfgCompany["ad_title"].(string),
					Atx: cfgCompany["ad_text"].(string),
					Aic: cfgCompany["ad_icon"].(string),
					Aig: cfgCompany["ad_image"].(string),
					Ccr: cfgCompany["company_country"].(string),
				}
				Campaigns = append(Campaigns, _Campany)

			}
		}

		if len(Campaigns) > 0 {

			rand.Seed(time.Now().Unix())
			n := rand.Int() % len(Campaigns)
			_creative := Campaigns[n]

			var timeDate int64
			now := time.Now()
			timeDate = now.Unix()

			linkData := LinkData{
				Link:   _creative.Cur,
				Cpc:    _creative.Cpr,
				Uid:    _creative.UID,
				Cid:    _creative.Cid,
				Cou:    country,
				Bro:    browser,
				Os:     platform,
				Sid:    sourceId,
				Date:   time.Unix(timeDate, 0).Format("2006-01-02"),
				Fresh:  ts.Freshness(timestamp),
				FeedId: feedId,
				Key:    uuid.New(),
				Click:  false,
			}

			jsonLink, _ := json.Marshal(linkData)

			var link = ""
			link = encrypt.Encrypt(string(jsonLink), config.Config["Crypto"].(string))
			link = config.Config["Click_Url"].(string) + "/click?data=" + link

			var creative = map[string]interface{}{
				"cpc":         _creative.Cpr - (_creative.Cpr * config.Config["revshare"].(float64) / 100),
				"ad_id":       _creative.Cid,
				"title":       _creative.Atl,
				"description": _creative.Atx,
				"icon":        config.Config["Media_Url"].(string) + "/" + _creative.Aic,
				"image":       config.Config["Media_Url"].(string) + "/" + _creative.Aig,
				"link":        link,
			}

			json, _ := json.Marshal(creative)

			waitGroup.Add(1)
			go addReq(linkData, &waitGroup, mongoClient)

			w.Write(json)
			w.WriteHeader(200)
			return
		} else {
			w.WriteHeader(204)
			return
		}

	})
	http.ListenAndServe(":9099", nil)
}

func addReq(data LinkData, waitGroup *sync.WaitGroup, client *mongo.Client) {
	defer waitGroup.Done()
	collection := client.Database(config.Config["mongo_database"].(string)).Collection(config.Config["mongo_collection"].(string))
	result, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
