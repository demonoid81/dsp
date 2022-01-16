package main

import (
	"context"
	"fmt"
	"github.com/demonoid81/dsp/auction/dsp"
	"github.com/demonoid81/dsp/auction/ssp"
	"github.com/demonoid81/dsp/events/kafkaMessage"
	"github.com/demonoid81/dsp/events/mongodb"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
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
	"github.com/demonoid81/dsp/events/utils"
	"github.com/google/uuid"
	"github.com/rs/cors"
)

type app struct {
	SSP         []dsp.SSP
	mongoClient *mongo.Client
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

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

var App *app

func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of get requests.",
	},
	[]string{"path"},
)

var totalRequestsByFeed = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_by_feed_total",
		Help: "Number of get requests by feed.",
	},
	[]string{"path", "feed"},
)

var responseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "response_status",
		Help: "Status of HTTP response",
	},
	[]string{"status"},
)

var httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "http_response_time_seconds",
	Help: "Duration of HTTP requests.",
}, []string{"path"})

func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()

		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
		rw := NewResponseWriter(w)
		next.ServeHTTP(rw, r)

		statusCode := rw.statusCode

		responseStatus.WithLabelValues(strconv.Itoa(statusCode)).Inc()
		totalRequests.WithLabelValues(path).Inc()

		timer.ObserveDuration()
	})
}

func init() {
	prometheus.Register(totalRequests)
	prometheus.Register(totalRequestsByFeed)
	prometheus.Register(responseStatus)
	prometheus.Register(httpDuration)
	App = &app{}
}

func main() {

	ctx := context.Background()

	runtime.GOMAXPROCS(runtime.NumCPU())

	var waitGroup sync.WaitGroup

	mongoClient, err := mongodb.NewClient()

	App.mongoClient = mongoClient

	if err != nil {
		fmt.Println(err)
		return
	}

	//err = App.uploadSSP(ctx)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	err = App.loadSSP(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	router := mux.NewRouter()
	router.Use(prometheusMiddleware)

	router.Path("/prometheus").Handler(promhttp.Handler())

	router.Path("/ssp").Handler(sspEvent(ctx, &waitGroup, mongoClient))

	router.Path("/subscribe").Handler(subscribe(ctx))

	router.Path("/click").Handler(click(ctx, &waitGroup, mongoClient))

	router.Path("/clickdsp").Handler(clickdsp(ctx))

	router.Path("/feed").Handler(ssp.Feed(ctx, App.SSP, &waitGroup, mongoClient, totalRequestsByFeed))

	router.Path("/stat").Handler(App.stat(ctx))

	router.Path("/ssp/get").Handler(App.getSSP(ctx))

	router.Path("/ssp/add").Handler(addSSP(ctx, mongoClient))
	router.Path("/ssp/update").Handler(addSSP(ctx, mongoClient))
	router.Path("/ssp/reload").Handler(App.reloadSSP(ctx))

	router.Path("/dsp/get").Handler(App.getDSP(ctx))
	router.Path("/dsp/add").Methods("POST").Handler(App.addDSP(ctx))
	router.Path("/dsp/update").Methods("POST").Handler(App.updateDSP(ctx))

	ui := UIHandler{staticFS: staticFiles, staticPath: "web/dist", indexPath: "index.html"}
	router.PathPrefix("/").Handler(ui)

	corsHandler := cors.Default().Handler(router)

	fmt.Println("Serving requests on port 9099")
	err = http.ListenAndServe(":9099", corsHandler)
	fmt.Println(err)
}

//func (app *app) uploadSSP(ctx context.Context) error {
//	collection := app.mongoClient.Database(config.Config["mongo_database"].(string)).Collection("ssp")
//	for _, ssp := range ssp.SSPData {
//		result, err := collection.InsertOne(ctx, ssp)
//		if err != nil {
//			fmt.Println(err)
//		}
//		fmt.Println(result)
//	}
//	return nil
//}

//func sspFeed(ctx context.Context, SSPData []dsp.SSP,  waitGroup *sync.WaitGroup, mongoClient *mongo.Client) http.HandlerFunc {
//	return ssp.Feed(ctx, app.SSP, waitGroup, mongoClient)
//}

func (app *app) reloadSSP(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := app.loadSSP(ctx)
		if err != nil {
			w.WriteHeader(503)
			return
		}
	}
}

func sspEvent(ctx context.Context, waitGroup *sync.WaitGroup, mongoClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rdb := redis.Client()

		params, err := url.PathUnescape(r.URL.RawQuery)
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

		browser := utils.GetBrowser(ua)
		platform := utils.GetPlatform(ua)

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
				Key:    uuid.New().String(),
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
			go addReq(ctx, linkData, waitGroup, mongoClient)

			w.Write(json)
			w.WriteHeader(200)
			return
		} else {
			w.WriteHeader(204)
			return
		}
	}
}

func subscribe(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Redirect(w, r, config.Config["Url_Redirect"].(string), 301)
		}

		bodyRequest := map[string]interface{}{}
		json.Unmarshal([]byte(body), &bodyRequest)

		var timeDate int64
		now := time.Now()
		timeDate = now.Unix()

		var subscribe = map[string]interface{}{}

		subscribe["subscriber_ip"] = utils.GetIP(r)
		subscribe["subscriber_user_agent"] = utils.GetUA(r)
		subscribe["subscriber_country"] = utils.GetCountry(subscribe["subscriber_ip"].(string))
		subscribe["subscriber_os"] = utils.GetOS(subscribe["subscriber_user_agent"].(string))
		subscribe["subscriber_browser"] = utils.GetBrowser(subscribe["subscriber_user_agent"].(string))
		subscribe["subscriber_date"] = time.Unix(timeDate, 0).Format("2006-01-02")
		subscribe["subscriber_last_send"] = 0
		subscribe["user_id"] = bodyRequest["user_id"]
		subscribe["stream_id"] = bodyRequest["stream_id"]
		subscribe["promo_id"] = bodyRequest["promo_id"]
		subscribe["subscriber_endpoint"] = bodyRequest["endpoint"]
		subscribe["subscriber_key"] = bodyRequest["key"]
		subscribe["subscriber_auth"] = bodyRequest["auth"]

		json, _ := json.Marshal(subscribe)

		go kafkaMessage.SendMessage(ctx, string(json), config.Config["Kafka"].(map[string]interface{})["subscribe"].(map[string]interface{}))
	}
}

func click(ctx context.Context, waitGroup *sync.WaitGroup, mongoClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keys := r.URL.Query()
		dataGet := keys.Get("data")

		fmt.Println(dataGet)

		if dataGet != "" {

			jsonData := encrypt.Decrypt(dataGet, config.Config["Crypto"].(string))

			data := make(map[string]interface{})
			json.Unmarshal([]byte(jsonData), &data)

			// data["shows"] = Random.E(100, 900)

			var link = data["link"].(string)

			if strings.Contains(link, "labyrinthads.com") {

				u, err := url.Parse(link)
				if err != nil {
					http.Redirect(w, r, config.Config["Url_Redirect"].(string), 302)
					return
				}
				fmt.Println(u.RawQuery)

				q, err := url.ParseQuery(u.RawQuery)
				if err != nil {
					http.Redirect(w, r, config.Config["Url_Redirect"].(string), 302)
					return
				}

				dataGet := q.Get("data")

				if dataGet != "" {
					jsonData := encrypt.Decrypt(dataGet, config.Config["Crypto"].(string))

					json.Unmarshal([]byte(jsonData), &data)
					link = data["link"].(string)

				} else {
					http.Redirect(w, r, config.Config["Url_Redirect"].(string), 302)
					return
				}
			}

			if strings.Contains(link, "{SOURCE_ID}") {
				link = strings.Replace(link, "{SOURCE_ID}", data["sid"].(string), -1)
			}
			if strings.Contains(link, "{CAMPAIGN_ID}") {
				link = strings.Replace(link, "{CAMPAIGN_ID}", data["cid"].(string), -1)
			}
			if strings.Contains(link, "{COST}") {
				link = strings.Replace(link, "{COST}", fmt.Sprintf("%f", data["cpc"]), -1)
			}
			if strings.Contains(link, "{COUNTRY}") {
				link = strings.Replace(link, "{COUNTRY}", data["cou"].(string), -1)
			}
			if strings.Contains(link, "{BROWSER}") {
				link = strings.Replace(link, "{BROWSER}", data["bro"].(string), -1)
			}
			if strings.Contains(link, "{OS}") {
				link = strings.Replace(link, "{OS}", data["os"].(string), -1)
			}
			_, fresh := data["fresh"]
			if fresh {
				if strings.Contains(link, "{FRESHNESS}") {
					link = strings.Replace(link, "{FRESHNESS}", data["fresh"].(string), -1)
				}
			}

			fmt.Println("feed_id:", data["feed_id"], " ", data["feed_id"].(string))

			if _, feedId := data["feed_id"]; feedId {
				if strings.Contains(link, "{FEED_ID}") {
					link = strings.Replace(link, "{FEED_ID}", data["feed_id"].(string), -1)
				}
			}

			data["link"] = link

			jsonKafka, _ := json.Marshal(data)

			if key, ok := data["uuid"]; ok {
				waitGroup.Add(1)
				go updateReq(ctx, key.(string), waitGroup, mongoClient)
			}

			totalRequestsByFeed.WithLabelValues("click", data["feed_id"].(string)).Inc()

			go kafkaMessage.SendMessage(ctx, string(jsonKafka), config.Config["Kafka"].(map[string]interface{})["click"].(map[string]interface{}))

			http.Redirect(w, r, data["link"].(string), 302)

		} else {
			http.Redirect(w, r, config.Config["Url_Redirect"].(string), 302)
		}
	}
}

func clickdsp(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keys := r.URL.Query()
		dataGet := keys.Get("data")

		if dataGet != "" {

			jsonData := encrypt.DecryptNode(dataGet, config.Config["Crypto"].(string))

			data := make(map[string]interface{})
			json.Unmarshal([]byte(jsonData), &data)

			//data["shows"] = Random.E(100, 900)

			var link = data["link"].(string)

			data["link"] = link

			jsonKafka, _ := json.Marshal(data)

			go kafkaMessage.SendMessage(ctx, string(jsonKafka), config.Config["Kafka"].(map[string]interface{})["clickdsp"].(map[string]interface{}))

			http.Redirect(w, r, data["link"].(string), 301)

		} else {
			http.Redirect(w, r, config.Config["Url_Redirect"].(string), 301)
		}
	}
}

func addReq(ctx context.Context, data LinkData, waitGroup *sync.WaitGroup, client *mongo.Client) {
	defer waitGroup.Done()
	collection := client.Database(config.Config["mongo_database"].(string)).Collection(config.Config["mongo_collection"].(string))
	result, err := collection.InsertOne(ctx, data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func updateReq(ctx context.Context, key string, waitGroup *sync.WaitGroup, client *mongo.Client) {
	defer waitGroup.Done()
	collection := client.Database(config.Config["mongo_database"].(string)).Collection(config.Config["mongo_collection"].(string))
	filter := bson.M{
		"uuid": bson.M{
			"$eq": key, // check if bool field has value of 'false'
		},
	}
	update := bson.M{
		"$set": bson.M{
			"click": true,
		},
	}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
