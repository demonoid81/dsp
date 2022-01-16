package ssp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/demonoid81/dsp/auction/dsp"
	"github.com/demonoid81/dsp/config"
	"github.com/demonoid81/dsp/events/encrypt"
	"github.com/demonoid81/dsp/events/utils"
	"github.com/prometheus/client_golang/prometheus"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

func Feed(ctx context.Context, SSPData []dsp.SSP, waitGroup *sync.WaitGroup, mongoClient *mongo.Client, counter *prometheus.CounterVec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		sspKey := r.FormValue("key")

		if sspKey == "" {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		idx := utils.Find(SSPData, func(value interface{}) bool {

			return value.(dsp.SSP).Key == sspKey
		})

		if idx < 0 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		country := utils.GetCountry(r.FormValue("ip"))

		data := dsp.ReqData{
			IP:      r.FormValue("ip"),
			UA:      r.FormValue("ua"),
			ID:      "",
			SID:     r.FormValue("id"), // to do в таблицу
			Time:    r.FormValue("time"),
			UID:     r.FormValue("uid"),
			Lang:    r.FormValue("lang"),
			TZ:      r.FormValue("tz"),
			Country: country,
		}

		creative, dataBase64, _ := dsp.Event(ctx, data, SSPData[idx], waitGroup, mongoClient)
		if creative.Status == 200 {

			w.Header().Set("Token", dataBase64)

			result := map[string]interface{}{}

			linkData := map[string]interface{}{
				"link":     creative.Link,
				"cpc":      creative.Cpc,
				"cpc_orig": creative.CpcOriginal,
				"dsp_id":   creative.DSPID,
				"dsp_name": creative.DSPName,
				"ssp_id":   creative.SSPID,
				"ssp_name": creative.SSPName,
				"ip":       r.FormValue("ip"),
				"country":  country,
				"clid":     rand.Intn(999999999-9999999) + 9999999,
				"id":       r.FormValue("id"),
				"uid":      r.FormValue("uid"),
			}

			jsonLink, _ := json.Marshal(linkData)

			var link = ""
			link = encrypt.Encrypt(string(jsonLink), config.Config["Crypto"].(string))
			link = config.Config["Click_Url"].(string) + "/click?data=" + link

			switch creative.SSPName {
			case "clickadu":
				cpc, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", creative.Cpc), 8)
				result = map[string]interface{}{
					"title":       creative.Title,
					"description": creative.Body,
					"price":       cpc,
					"image":       creative.Image,
					"icon":        creative.Icon,
					"url":         link,
				}
			case "adskeeper":
				cpc, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", creative.Cpc), 8)
				result = map[string]interface{}{
					"text":        creative.Body,
					"title":       creative.Title,
					"cpc":         cpc,
					"click_url":   link,
					"image_url":   creative.Image,
					"icon_url":    creative.Icon,
					"campaign_id": creative.SSPID,
					"category":    "1",
					"id":          creative.ID,
				}
			default:
				cpc, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", creative.Cpc), 8)
				result = map[string]interface{}{
					"id":          creative.ID,
					"title":       creative.Title,
					"description": creative.Body,
					"icon":        creative.Icon,
					"image":       creative.Image,
					"url":         link,
					"bid":         cpc,
				}
			}

			counter.WithLabelValues("feed", fmt.Sprintf("$d", creative.SSPID)).Inc()

			res, err := json.Marshal(result)
			if err != nil {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			w.Write(res)
			w.WriteHeader(http.StatusOK)
			return

		}
		w.WriteHeader(http.StatusNoContent)
	}
}