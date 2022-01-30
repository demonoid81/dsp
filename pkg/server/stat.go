package server

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)



func (s *Server) stat(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startDate, err := time.Parse("2006-01-02", r.FormValue("start"))
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(503)
			return
		}
		endDate, err := time.Parse("2006-01-02", r.FormValue("end"))
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(503)
			return
		}

		feedID := r.FormValue("feed_id")

		type status struct {
			Date  string  `json:"date"`
			Shows int64   `json:"shows"`
			Click int64   `json:"click"`
			Rate  float64 `json:"rate"`
			CPC   float64 `json:"cpc"`
			CTR   float64 `json:"ctr"`
		}

		days := endDate.Sub(startDate).Hours() / 24
		fmt.Println(days)
		var statuses []status
		collection := s.mongo.MongoClient.Database(s.cfg.MongoDatabase).Collection("requests")
		for i := 0; i <= int(days); i++ {

			date := startDate.Add(time.Hour * 24 * time.Duration(i)).Format("2006-01-02")

			fmt.Println(date)

			filter := bson.M{
				"date": bson.M{
					"$eq": date, // check if bool field has value of 'false'
				},
			}
			if feedID != "" {
				filter = bson.M{
					"date":    bson.M{"$eq": date},
					"feed_id": bson.M{"$eq": feedID},
				}
			}

			shows, err := collection.CountDocuments(ctx, filter)
			if err != nil {
				w.WriteHeader(503)
			}

			filter = bson.M{
				"date":  bson.M{"$eq": date},
				"click": bson.M{"$eq": true},
			}
			if feedID != "" {
				filter = bson.M{
					"date":    bson.M{"$eq": date},
					"click":   bson.M{"$eq": true},
					"feed_id": bson.M{"$eq": feedID},
				}
			}

			clicks, err := collection.CountDocuments(ctx, filter)
			if err != nil {
				w.WriteHeader(503)
			}

			curStat := status{
				Date:  date,
				Shows: shows,
				Click: clicks,
			}

			statuses = append(statuses, curStat)
		}

		data, err := json.Marshal(statuses)
		w.Write(data)
		w.WriteHeader(200)
	}
}
