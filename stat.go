package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/demonoid81/dsp/config"
	"github.com/demonoid81/dsp/json2table"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

type LData struct {
	Country string `json:"cou" bson:"country"`
	Browser string `json:"bro" bson:"browser"`
	Os      string `json:"os" bson:"os"`
	Sid     string `json:"sid" bson:"sid"`
	Date    string `json:"date" bson:"date"`
	FeedId  string `json:"feed_id" bson:"feed_id"`
	ReqFeed int64  `json:"req_feed" bson:"req_feed"`
	Clicks  int64  `json:"clicks" bson:"clicks"`
}

func (app *app) RebuldStat(ctx context.Context) error {
	statCollection := app.mongoClient.Database(config.Config["mongo_database"].(string)).Collection("statistics")
	requestsCollection := app.mongoClient.Database(config.Config["mongo_database"].(string)).Collection("requests")

	cur, err := requestsCollection.Find(ctx, bson.D{{}})
	if err != nil {
		return err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		//Create a value into which the single document can be decoded
		var elem LinkData
		err := cur.Decode(&elem)
		if err != nil {
			return err
		}
		filter := bson.M{
			"date":    bson.M{"$eq": elem.Date},
			"feed_id": bson.M{"$eq": elem.FeedId},
			"country": bson.M{"$eq": elem.Cou},
			"browser": bson.M{"$eq": elem.Bro},
			"os":      bson.M{"$eq": elem.Os},
			"sid":     bson.M{"$eq": elem.Sid},
		}
		var data LData
		if err := statCollection.FindOne(ctx, filter).Decode(data); err != nil {
			if err == mongo.ErrNoDocuments {
				data := LData{
					Country: elem.Cou,
					Browser: elem.Bro,
					Os:      elem.Os,
					Sid:     elem.Sid,
					Date:    elem.Date,
					FeedId:  elem.FeedId,
					ReqFeed: 1,
					Clicks:  0,
				}
				if elem.Click {
					data.Clicks = 1
				}
				result, err := statCollection.InsertOne(ctx, data)
				if err != nil {
					return err
				}
				fmt.Println(result)
			} else {
				return err
			}
			if elem.Click {
				data.Clicks = data.Clicks + 1
			}
			update := bson.M{
				"$set": bson.M{
					"req_feed": data.ReqFeed + 1,
					"clicks":   data.Clicks,
				},
			}
			result, err := statCollection.UpdateOne(ctx, filter, update)
			if err != nil {
				return err
			}
			fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
		}

	}
	return nil
}

func (app *app) stat(ctx context.Context) http.HandlerFunc {
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
		collection := app.mongoClient.Database(config.Config["mongo_database"].(string)).Collection(config.Config["mongo_collection"].(string))
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

		fmt.Println(data)

		_, html := json2table.JSON2HtmlTable(string(data), nil, nil)

		w.Write([]byte(html))
		w.WriteHeader(200)
	}
}
