package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/demonoid81/dsp/config"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type DSP struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Endpoint string `json:"endpoint"`
	Type     string `json:"type"`
	QPS      int    `json:"qps"`
}

func (app *app) getDSP(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ssp []DSP

		collection := app.mongoClient.Database(config.Config["mongo_database"].(string)).Collection("dsp")
		cur, err := collection.Find(ctx, bson.D{{}})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer cur.Close(ctx)

		for cur.Next(ctx) {
			//Create a value into which the single document can be decoded
			var elem DSP
			err := cur.Decode(&elem)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			ssp = append(ssp, elem)

		}
		if err := cur.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Printf("Found multiple documents: %+v\n", ssp)
		res, err := json.Marshal(ssp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(res)
		w.WriteHeader(http.StatusOK)
	}
}

func (app *app) addDSP(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dsp DSP
		err := json.NewDecoder(r.Body).Decode(&dsp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(dsp)
		collection := app.mongoClient.Database(config.Config["mongo_database"].(string)).Collection("dsp")
		result, err := collection.InsertOne(ctx, dsp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(result)
		w.WriteHeader(http.StatusOK)
	}
}

