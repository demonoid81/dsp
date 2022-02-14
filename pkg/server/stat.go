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

		filter := r.FormValue("filter")

		type status struct {
			Metrica string  `json:"metrica" bson:"metrica"`
			ReqFeed int64   `json:"req_feed" bson:"req_feed"`
			Clicks  int64   `json:"clicks" bson:"clicks"`
			Rate    float64 `json:"rate" bson:"rate"`
			CPC     float64 `json:"cpc" bson:"cpc"`
			CTR     float64 `json:"ctr" bson:"ctr"`
		}

		days := endDate.Sub(startDate).Hours() / 24
		var statuses []status

		collection := s.mongo.MongoClient.Database(s.cfg.MongoDatabase).Collection("statistics")

		var dates []string
		for i := 0; i <= int(days); i++ {
			date := startDate.Add(time.Hour * 24 * time.Duration(i)).Format("2006-01-02")
			dates = append(dates, date)
		}

		matchStage := bson.M{
			"$match": bson.M{
				"date": bson.M{
					"$in": dates,
				}}}
		groupStage := bson.M{
			"$group": bson.M{
				"_id": fmt.Sprintf("$%s", filter),
				"total_click": bson.M{
					"$sum": "$clicks",
				},
				"total_req_feed": bson.M{
					"$sum": "$req_feed",
				},
				"total_rate": bson.M{
					"$sum": "$rate",
				},
			},
		}
		projectStage := bson.M{
			"$project": bson.M{
				"metrica":  "$_id",
				"clicks":   "$total_click",
				"req_feed": "$total_req_feed",
				"rate":     "$total_rate",
			},
		}

		cursor, err := collection.Aggregate(ctx, []bson.M{matchStage, groupStage, projectStage})
		if err != nil {
			panic(err)
		}

		if err = cursor.All(ctx, &statuses); err != nil {
			w.WriteHeader(503)
		}

		data, err := json.Marshal(statuses)
		w.Write(data)
		w.WriteHeader(200)
	}
}
