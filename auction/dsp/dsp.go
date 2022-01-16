package dsp

import (
	"context"
	"encoding/base64"
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
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type DSPCfg struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Endpoint string `json:"endpoint"`
	Type     string `json:"type"`
	QPS      int64  `json:"qps"`
}

var DSPData = []DSPCfg{{
	ID:       102,
	Name:     "labyrinthads",
	Endpoint: "http://dsp.labyrinthads.com/ssp?key=adexchange&category=1&push_type=classic",
	Type:     "mainstream",
	QPS:      10000,
}, {
	ID:       103,
	Name:     "labyrinthads",
	Endpoint: "http://dsp.labyrinthads.com/ssp?key=adexchange&category=1&push_type=inpage",
	Type:     "mainstream",
	QPS:      10000,
},
}

type ReqData struct {
	IP      string
	UA      string
	ID      string
	SID     string
	Time    string
	UID     string
	Lang    string
	TZ      string
	Country string
}

type Creative struct {
	Body        string  `json:"body"`
	Cpc         float64 `json:"cpc"`
	CpcOriginal float64 `json:"cpc_original"`
	Icon        string  `json:"icon"`
	ID          string  `json:"id"`
	Image       string  `json:"image"`
	Link        string  `json:"link"`
	Status      int     `json:"status"`
	Title       string  `json:"title"`
	DSPID       int     `json:"dsp_id"`
	DSPName     string  `json:"dsp_name"`
	SSPID       int     `json:"ssp_id"`
	SSPName     string  `json:"ssp_name"`
	Timestamp   int     `json:"timestamp"`
}

var Affiliates = map[string]func(ctx context.Context, data ReqData, config DataDSP, waitGroup *sync.WaitGroup, mongoClient *mongo.Client) Creative{
	"labyrinthads": Get,
}

type single struct {
	mu     sync.Mutex
	values map[int]string
}

var counters = single{
	values: make(map[int]string),
}

func (s *single) Get(key int) int64 {
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

func (s *single) Set(key int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	vals := strings.Split(s.values[key], "_")
	number, _ := strconv.ParseInt(vals[1], 10, 32)
	count := number + 1
	nowDate := time.Now()
	timestamp := nowDate.Unix()
	s.values[key] = strconv.FormatInt(timestamp, 10) + "_" + strconv.FormatInt(count, 10)
}

var Creatives []Creative

type DSP struct {
	ID                int `json:"id"`
	Profit            float64
	SourceIdBlacklist []string `json:"source_id_blacklist"`
	CountryBlacklist  []string `json:"country_blacklist"`
	CountryWhitelist  []string `json:"country_whitelist"`
}

type SSP struct {
	Key  string `json:"key"`
	Name string `json:"ssp_name"`
	ID   int    `json:"ssp_id"`
	DSP  []DSP  `json:"dsp"`
	Type string `json:"type"`
}

type DataDSP struct {
	DSPID    int
	DSPName  string
	Profit   float64
	Endpoint string
	SSPID    int
	SSPName  string
}

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
	Key    string  `json:"uuid" bson:"uuid"`
	Link   string  `json:"link" bson:"link"`
	Cpc    float64 `json:"cpc" bson:"cpc"`
	Uid    float64 `json:"uid" bson:"uid"`
	Cid    string  `json:"cid" bson:"cid"`
	Cou    string  `json:"cou" bson:"cou"`
	Bro    string  `json:"bro" bson:"bro"`
	Os     string  `json:"os" bson:"os"`
	Sid    string  `json:"sid" bson:"sid"`
	Date   string  `json:"date" bson:"date"`
	Fresh  string  `json:"fresh" bson:"fresh"`
	FeedId string  `json:"feed_id" bson:"feed_id"`
	Click  bool    `json:"-" bson:"click"`
}

func Event(ctx context.Context, data ReqData, sspData SSP, waitGroup *sync.WaitGroup, mongoClient *mongo.Client) (*Creative, string, error) {

	var creatives []Creative

	for _, dsp := range sspData.DSP {
		idx := utils.Find(DSPData, func(value interface{}) bool {
			return value.(DSPCfg).ID == dsp.ID
		})

		if idx < 0 {
			return nil, "", fmt.Errorf("dsp not found")
		}

		var cfg = DataDSP{
			DSPID:    dsp.ID,
			DSPName:  DSPData[idx].Name,
			Profit:   dsp.Profit,
			Endpoint: DSPData[idx].Endpoint,
			SSPID:    sspData.ID,
			SSPName:  sspData.Name,
		}

		if (!utils.ContainsInArray(dsp.SourceIdBlacklist, data.SID) &&
			!utils.ContainsInArray(dsp.CountryBlacklist, data.Country) &&
			(len(dsp.CountryWhitelist) == 0 || utils.ContainsInArray(dsp.CountryWhitelist, data.Country))) &&
			counters.Get(dsp.ID) <= DSPData[idx].QPS/2 {

			counters.Set(dsp.ID)

			data.ID = fmt.Sprintf("%d%d%s", dsp.ID, sspData.ID, data.SID)

			res := Affiliates[DSPData[idx].Name](ctx, data, cfg, waitGroup, mongoClient)
			creatives = append(creatives, res)
		}
	}

	sort.Slice(creatives, func(i, j int) bool {
		return creatives[i].Cpc > creatives[j].Cpc
	})

	returnCreative := Creative{}

	if len(creatives) > 0 {
		returnCreative = creatives[0]
	}

	var headerCreatives []map[string]interface{}

	for index, crKafka := range creatives {
		cr_ := map[string]interface{}{
			"status": crKafka.Status,
			"d_id":   crKafka.DSPID,
		}
		if crKafka.Status == 200 {
			if index == 0 {
				cr_["w"] = 1
				cr_["cpc"] = crKafka.CpcOriginal
			}
		}
		headerCreatives = append(headerCreatives, cr_)
	}

	jsonHeader, _ := json.Marshal(headerCreatives)
	base64Header := base64.StdEncoding.EncodeToString(jsonHeader)

	return &returnCreative, base64Header, nil

}

func Get(ctx context.Context, data ReqData, cfg DataDSP, waitGroup *sync.WaitGroup, mongoClient *mongo.Client) Creative {

	rdb := redis.Client()

	country := data.Country
	sourceId := data.ID
	category := "1"
	timestamp := data.Time
	pushType := "classic"
	feedId := cfg.SSPID

	ua := strings.ReplaceAll(data.UA, "+", " ")
	browser := utils.GetBrowser(ua)
	platform := utils.GetPlatform(ua)

	var redisKey = fmt.Sprintf("%s-%s-%s-%s-%s-%d", browser, platform, country, category, pushType, feedId)

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
				return Creative{
					Status:  204,
					DSPID:   cfg.DSPID,
					DSPName: cfg.DSPName,
					SSPID:   cfg.SSPID,
					SSPName: cfg.SSPName,
				}
			}

		} else {
			return Creative{
				Status:  204,
				DSPID:   cfg.DSPID,
				DSPName: cfg.DSPName,
				SSPID:   cfg.SSPID,
				SSPName: cfg.SSPName,
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
			(len(whitelist) == 0 || inArray.FindString(whitelist, sourceId) == true) &&
			inArray.FindString(blacklistFeed, fmt.Sprintf("%d", feedId)) == false &&
			(len(whitelistFeed) == 0 || inArray.FindString(whitelistFeed, fmt.Sprintf("%d", feedId)) == true) {

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
			FeedId: fmt.Sprintf("%d", feedId),
			Key:    uuid.New().String(),
			Click:  false,
		}

		jsonLink, _ := json.Marshal(linkData)

		var link = ""
		link = encrypt.Encrypt(string(jsonLink), config.Config["Crypto"].(string))
		link = config.Config["Click_Url"].(string) + "/click?data=" + link

		waitGroup.Add(1)
		go mongodb.AddReq(ctx, linkData, waitGroup, mongoClient)

		cpc := _creative.Cpr - (_creative.Cpr * config.Config["revshare"].(float64) / 100)

		return Creative{
			Status:      200,
			ID:          _creative.Cid,
			Title:       _creative.Atl,
			Body:        _creative.Atx,
			Icon:        config.Config["Media_Url"].(string) + "/" + _creative.Aic,
			Image:       config.Config["Media_Url"].(string) + "/" + _creative.Aig,
			Link:        link,
			Cpc:         cpc - (cpc * cfg.Profit / 100),
			CpcOriginal: cpc,
			DSPID:       cfg.DSPID,
			DSPName:     cfg.DSPName,
			SSPID:       cfg.SSPID,
			SSPName:     cfg.SSPName,
		}
	} else {
		return Creative{
			Status:  499,
			DSPID:   cfg.DSPID,
			DSPName: cfg.DSPName,
			SSPID:   cfg.SSPID,
			SSPName: cfg.SSPName,
		}
	}
}
