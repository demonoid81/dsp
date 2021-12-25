package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/demonoid81/dsp/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type DSP struct {
	DSPID             string   `json:"dsp_id"`
	Profit            float64  `json:"profit"`
	SourceIdBlacklist []string `json:"source_id_blacklist"`
	CountryBlacklist  []string `json:"country_blacklist"`
	CountryWhitelist  []string `json:"country_whitelist"`
	Type string `json:"type"`
}

type SSP struct {
	Key   string `json:"key"`
	Name  string `json:"ssp_name"`
	SSPID string `json:"ssp_id"`
	DSP   DSP    `json:"dsp"`
	Type string `json:"type"`
}

func (app *app) loadSSP(ctx context.Context) error {
	var ssp []SSP

	collection := app.mongoClient.Database(config.Config["mongo_database"].(string)).Collection("ssp")
	cur, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		return err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		//Create a value into which the single document can be decoded
		var elem SSP
		err := cur.Decode(&elem)
		if err != nil {
			return err
		}

		ssp = append(ssp, elem)

	}
	if err := cur.Err(); err != nil {
		return err
	}
	fmt.Printf("Found multiple documents: %+v\n", ssp)
	return nil
}

func (app *app) getSSP(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	var ssp []SSP

	collection := app.mongoClient.Database(config.Config["mongo_database"].(string)).Collection("ssp")
	cur, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		//Create a value into which the single document can be decoded
		var elem SSP
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

func addSSP(ctx context.Context, client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ssp SSP
		err := json.NewDecoder(r.Body).Decode(&ssp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(ssp)
		collection := client.Database(config.Config["mongo_database"].(string)).Collection("ssp")
		result, err := collection.InsertOne(ctx, ssp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(result)
		w.WriteHeader(http.StatusOK)
	}
}
