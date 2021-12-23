package labyrinthads

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/demonoid81/dsp/config"
	"github.com/demonoid81/dsp/events/encrypt"
	"github.com/demonoid81/dsp/events/inArray"
	"github.com/demonoid81/dsp/events/mongodb"
	"github.com/demonoid81/dsp/events/postgres"
	"github.com/demonoid81/dsp/events/redis"
	ts "github.com/demonoid81/dsp/events/timestamp"
	"github.com/demonoid81/dsp/events/utils"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
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
	Key   string `json:"uuid" bson:"uuid"`
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
	Click  bool      `json:"-" bson:"click"`
}

type Creatives struct {
	AdID        string  `json:"ad_id"`
	Cpc         float64 `json:"cpc"`
	Description string  `json:"description"`
	Icon        string  `json:"icon"`
	Image       string  `json:"image"`
	Link        string  `json:"link"`
	Title       string  `json:"title"`
}

func Get(ctx context.Context, data map[string]interface{}, c chan map[string]interface{}, Config map[string]interface{}, waitGroup *sync.WaitGroup, mongoClient *mongo.Client) {

	rdb := redis.Client()

	country := data["country"].(string)
	sourceId := data["id"].(string)
	category := "1"
	timestamp := data["time"].(string)
	pushType := "classic"
	feedId := Config["ssp_id"].(string)

	ua := strings.ReplaceAll(data["ua"].(string), "+", " ")
	browser := utils.GetBrowser(ua)
	platform := utils.GetPlatform(ua)

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
				c <- map[string]interface{}{
					"status": 204,
					"dsp_id": Config["dsp_id"],
					"dsp_name": Config["dsp_name"],
					"ssp_id": Config["ssp_id"],
					"ssp_name": Config["ssp_name"],
				}
				return
			}

		} else {
			c <- map[string]interface{}{
				"status": 204,
				"dsp_id": Config["dsp_id"],
				"dsp_name": Config["dsp_name"],
				"ssp_id": Config["ssp_id"],
				"ssp_name": Config["ssp_name"],
			}
		}
	}

	conn := rdb.Get()
	conn.Close()

	var campaignsMap []map[string]interface{}
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
			Key:   uuid.New().String(),
			Click:  false,
		}

		jsonLink, _ := json.Marshal(linkData)

		fmt.Println(jsonLink)

		var link = ""
		link = encrypt.Encrypt(string(jsonLink), config.Config["Crypto"].(string))
		link = config.Config["Click_Url"].(string) + "/click?data=" + link

		waitGroup.Add(1)
		go  mongodb.AddReq(ctx, linkData, waitGroup, mongoClient)

		cpc := _creative.Cpr - (_creative.Cpr * config.Config["revshare"].(float64) / 100)

		c <- map[string]interface{}{
			"status": 200,
			"id": _creative.Cid,
			"title": _creative.Atl,
			"body": _creative.Atx,
			"icon": config.Config["Media_Url"].(string) + "/" + _creative.Aic,
			"image": config.Config["Media_Url"].(string) + "/" + _creative.Aig,
			"link": link,
			"cpc": cpc - (cpc * Config["profit"].(float64) / 100),
			"cpc_original": cpc,
			"dsp_id": Config["dsp_id"],
			"dsp_name": Config["dsp_name"],
			"ssp_id": Config["ssp_id"],
			"ssp_name": Config["ssp_name"],
		}
	} else {
		c <- map[string]interface{}{
			"status": 499,
			"dsp_id": Config["dsp_id"],
			"dsp_name": Config["dsp_name"],
			"ssp_id": Config["ssp_id"],
			"ssp_name": Config["ssp_name"],
		}
		return
	}
}