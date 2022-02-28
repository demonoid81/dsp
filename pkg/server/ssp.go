package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/demonoid81/dsp/auction/dsp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strconv"
)

func (s *Server) getSSP(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ssp []dsp.SSP

		collection := s.mongo.MongoClient.Database(s.cfg.MongoDatabase).Collection("ssp")
		cur, err := collection.Find(ctx, bson.D{{}})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer cur.Close(ctx)

		for cur.Next(ctx) {
			//Create a value into which the single document can be decoded
			var elem dsp.SSP
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

// Добавлении  SSP
func (s *Server) addSSP(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ssp dsp.SSP
		err := json.NewDecoder(r.Body).Decode(&ssp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		opts := options.Update().SetUpsert(true)

		query := bson.M{"id": bson.M{"$eq": ssp.ID}}

		update := bson.M{"$set": bson.M{
			"id":       ssp.ID,
			"name":     ssp.Name,
			"type": 	ssp.Type,
			"key":     	ssp.Key,
			"dsp":      ssp.DSP,
		}}

		collection := s.mongo.MongoClient.Database(s.cfg.MongoDatabase).Collection("ssp")
		result, err := collection.UpdateOne(ctx, query, update, opts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(result)
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) deleteSSP(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("id")

		fmt.Println(id)
		i, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		collection := s.mongo.MongoClient.Database(s.cfg.MongoDatabase).Collection("ssp")

		result, err := collection.DeleteOne(ctx, bson.M{"id": bson.M{"$eq": i}})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)
		w.WriteHeader(http.StatusOK)
	}
}

// Перегрузка при изменении задания
func (s *Server) reloadSSP(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
