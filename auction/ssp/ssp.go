package ssp

import (
	"context"
	"github.com/demonoid81/dsp/config"
	"github.com/demonoid81/dsp/events/utils"
	"net/http"
)

func feed(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx.Response.Header.Set("Content-Type", "application/json")

		ssp := string(r.FormValue("key"))
		cfg, err := config.Config["SSP"].(map[string]interface{})[ssp]

		if !err {
			w.WriteHeader(403)
			return
		}

		country := utils.GetCountry(string(r.FormValue("ip")))

		data := map[string]interface{}{
			"ip": string(r.FormValue("ip")),
			"ua": string(r.FormValue("ua")),
			"id": "",
			"sid": string(r.FormValue("id")),
			"time": string(r.FormValue("time")),
			"uid": string(r.FormValue("uid")),
			"lang": string(r.FormValue("lang")),
			"tz": string(r.FormValue("tz")),
			"country": country,
		}

		creative, data_base64 := edsp(data, cfg.(map[string]interface{}))

	}
}
