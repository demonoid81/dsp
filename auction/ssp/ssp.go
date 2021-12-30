package ssp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/demonoid81/dsp/auction/dsp"
	"github.com/demonoid81/dsp/config"
	"github.com/demonoid81/dsp/events/encrypt"
	"github.com/demonoid81/dsp/events/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)



func Feed(ctx context.Context, SSPData []dsp.SSP, waitGroup *sync.WaitGroup, mongoClient *mongo.Client) http.HandlerFunc {
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
			resultMultiple := []map[string]interface{}{}

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
				res := map[string]interface{}{
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
				resultMultiple = append(resultMultiple, res)
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
			if len(resultMultiple) > 0 {
				res, err := json.Marshal(resultMultiple)
				if err != nil {
					w.WriteHeader(http.StatusNoContent)
					return
				}
				w.Write(res)
				w.WriteHeader(http.StatusOK)
				return
			} else {
				res, err := json.Marshal(result)
				if err != nil {
					w.WriteHeader(http.StatusNoContent)
					return
				}
				w.Write(res)
				w.WriteHeader(http.StatusOK)
				return
			}
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

//var SSPData = []dsp.SSP{
//	{
//		Key:  "YxXs1HSVgDRRC9T",
//		Name: "Mgid",
//		ID:   102,
//		DSP: []dsp.DSP{
//			{
//				ID:                102,
//				Profit:            0.00,
//				SourceIdBlacklist: []string{},
//				CountryBlacklist:  []string{},
//				CountryWhitelist:  []string{},
//			},
//		},
//	},
//	{
//		Key:  "xBXhzjjhhYDFU7r",
//		Name: "test",
//		ID:   101,
//		DSP: []dsp.DSP{{
//			ID:                102,
//			Profit:            0.00,
//			SourceIdBlacklist: []string{"102101back_block_lp_2087", "102101ramos_w10_1304_us_edge_new", "102101lux_w10_0303-dagger-2033", "102101ramos_w10_0605_us_edge", "102101terame_w10_1609_edge_486231", "102101ramos_w10_1608_topiced", "102101ramos_w10_0907_us_edge", "102101ceca_w10_1608_kasa", "102101yue_w10_1208_manunited"},
//			CountryBlacklist:  []string{},
//			CountryWhitelist:  []string{},
//		}},
//	},
//	{
//		Key:  "aeW4bHWkaGIMl6N",
//		Name: "Clickadilla_In-Page",
//		ID:   111,
//		DSP: []dsp.DSP{
//			{
//				ID:                103,
//				Profit:            0.00,
//				SourceIdBlacklist: []string{},
//				CountryBlacklist:  []string{},
//				CountryWhitelist:  []string{},
//			},
//		},
//	},
//	{
//		Key:  "qjHLq1RZXU3tk8j",
//		Name: "Clickadilla_In-Page",
//		ID:   112,
//		DSP: []dsp.DSP{
//			{
//				ID:                103,
//				Profit:            0.00,
//				SourceIdBlacklist: []string{},
//				CountryBlacklist:  []string{},
//				CountryWhitelist:  []string{},
//			},
//		},
//	},
//	{
//		Key:  "wKyozxnWWgdHECi",
//		Name: "Clickadilla_In-Page",
//		ID:   113,
//		DSP: []dsp.DSP{
//			{
//				ID:                103,
//				Profit:            0.00,
//				SourceIdBlacklist: []string{},
//				CountryBlacklist:  []string{},
//				CountryWhitelist:  []string{},
//			},
//		},
//	},
//	{
//		Key:  "STx9xtDhW0825P8",
//		Name: "Clickadilla_In-Page",
//		ID:   114,
//		DSP: []dsp.DSP{
//			{
//				ID:                103,
//				Profit:            0.00,
//				SourceIdBlacklist: []string{},
//				CountryBlacklist:  []string{},
//				CountryWhitelist:  []string{},
//			},
//		},
//	},
//	{
//		Key:  "lz5ouNvlSnlZOIf",
//		Name: "DaoPush",
//		ID:   115,
//		DSP: []dsp.DSP{
//			{
//				ID:                102,
//				Profit:            0.00,
//				SourceIdBlacklist: []string{},
//				CountryBlacklist:  []string{},
//				CountryWhitelist:  []string{},
//			},
//		},
//	},
//	{
//		Key:  "UvEH4RJTsa2oOpx",
//		Name: "DaoPush_Classic",
//		ID:   116,
//		DSP: []dsp.DSP{
//			{
//				ID:                102,
//				Profit:            0.00,
//				SourceIdBlacklist: []string{},
//				CountryBlacklist:  []string{},
//				CountryWhitelist:  []string{},
//			},
//		},
//	},
//	{
//		Key:  "NaENhvugH5WdQ1y",
//		Name: "DaoPush_In-Page",
//		ID:   117,
//		DSP: []dsp.DSP{
//			{
//				ID:                103,
//				Profit:            0.00,
//				SourceIdBlacklist: []string{},
//				CountryBlacklist:  []string{},
//				CountryWhitelist:  []string{},
//			},
//		},
//	},
//	{
//		Key:  "ngGnozEJ3syxUEp",
//		Name: "DaoPush_In-Page",
//		ID:   118,
//		DSP: []dsp.DSP{
//			{
//				ID:                103,
//				Profit:            0.00,
//				SourceIdBlacklist: []string{},
//				CountryBlacklist:  []string{},
//				CountryWhitelist:  []string{},
//			},
//		},
//	},
//	{
//		Key:  "wQdvtTF2b6fs2p2",
//		Name: "RexRTB",
//		ID:   103,
//		DSP: []dsp.DSP{
//			{
//				ID:                102,
//				Profit:            0.00,
//				SourceIdBlacklist: []string{},
//				CountryBlacklist:  []string{},
//				CountryWhitelist:  []string{},
//			},
//		},
//	},
//	{
//		Key:  "PJtnhleyzNYAdGk",
//		Name: "RiverTraffic_Classic",
//		ID:   104,
//		DSP: []dsp.DSP{
//			{
//				ID:                102,
//				Profit:            0.00,
//				SourceIdBlacklist: []string{},
//				CountryBlacklist:  []string{},
//				CountryWhitelist:  []string{},
//			},
//		},
//	},
//	{
//		Key:  "z8GBO0tCQwTw2bt",
//		Name: "RiverTraffic_In-Page",
//		ID:   105,
//		DSP: []dsp.DSP{
//			{
//				ID:                103,
//				Profit:            0.00,
//				SourceIdBlacklist: []string{},
//				CountryBlacklist:  []string{},
//				CountryWhitelist:  []string{},
//			},
//		},
//	},
//	{
//		Key:  "wlFmvmEl5spoPxk",
//		Name: "RiverTraffic_In-Page",
//		ID:   106,
//		DSP: []dsp.DSP{
//			{
//				ID:                103,
//				Profit:            0.00,
//				SourceIdBlacklist: []string{},
//				CountryBlacklist:  []string{},
//				CountryWhitelist:  []string{},
//			},
//		},
//	},
//	{
//		Key:  "FqYgmeGAJG9nJz6",
//		Name: "Clickadilla_In-Page",
//		ID:   109,
//		DSP: []dsp.DSP{
//			{
//				ID:                103,
//				Profit:            0.00,
//				SourceIdBlacklist: []string{},
//				CountryBlacklist:  []string{},
//				CountryWhitelist:  []string{},
//			},
//		},
//	},
//	{
//		Key:  "vv8VZtulWS5KvvC",
//		Name: "Clickadilla_In-Page",
//		ID:   110,
//		DSP: []dsp.DSP{
//			{
//				ID:                103,
//				Profit:            0.00,
//				SourceIdBlacklist: []string{},
//				CountryBlacklist:  []string{},
//				CountryWhitelist:  []string{},
//			},
//		},
//	},
//	{
//		Key:  "s9u1H1aJuJgYWJc",
//		Name: "Admaven_Adult_push",
//		ID:   119,
//		DSP: []dsp.DSP{
//			{
//				ID:                102,
//				Profit:            0.00,
//				SourceIdBlacklist: []string{},
//				CountryBlacklist:  []string{},
//				CountryWhitelist:  []string{},
//			},
//		},
//	},
//	{
//		Key:  "O29B7nMFGE6fFNv",
//		Name: "Admaven_MNSTRM_push",
//		ID:   120,
//		DSP: []dsp.DSP{
//			{
//				ID:                102,
//				Profit:            0.00,
//				SourceIdBlacklist: []string{},
//				CountryBlacklist:  []string{},
//				CountryWhitelist:  []string{},
//			},
//		},
//	},
//	{
//		Key:  "uiUIGofvdGxzmA4",
//		Name: "Clickadilla_Classic",
//		ID:   107,
//		DSP: []dsp.DSP{
//			{
//				ID:                102,
//				Profit:            0.00,
//				SourceIdBlacklist: []string{"102108305275"},
//				CountryBlacklist:  []string{},
//				CountryWhitelist:  []string{},
//			},
//		},
//	},
//	{
//		Key:  "2S46Snsmc1IaISv",
//		Name: "Clickadilla_Classic",
//		ID:   108,
//		DSP: []dsp.DSP{
//			{
//				ID:                102,
//				Profit:            0.00,
//				SourceIdBlacklist: []string{"102108305275"},
//				CountryBlacklist:  []string{},
//				CountryWhitelist:  []string{},
//			},
//		},
//	},
//	{
//		Key:  "nimbHWu4CUtKUPP",
//		Name: "Clickadilla_Classic",
//		ID:   121,
//		DSP: []dsp.DSP{
//			{
//				ID:                102,
//				Profit:            0.00,
//				SourceIdBlacklist: []string{"102108305275"},
//				CountryBlacklist:  []string{},
//				CountryWhitelist:  []string{},
//			},
//		},
//	},
//	{
//		Key:  "4uM4MVUnQtV8EPv",
//		Name: "Clickadilla_Classic",
//		ID:   122,
//		DSP: []dsp.DSP{
//			{
//				ID:                102,
//				Profit:            0.00,
//				SourceIdBlacklist: []string{"102108305275"},
//				CountryBlacklist:  []string{},
//				CountryWhitelist:  []string{},
//			},
//		},
//	},
//}
