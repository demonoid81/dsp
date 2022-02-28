package server

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strconv"
)

//
import (
	"context"
	"encoding/json"
	"github.com/demonoid81/dsp/auction/dsp"
)

//
func (s *Server) getDSP(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var elements []dsp.DSPCfg

		collection := s.mongo.MongoClient.Database(s.cfg.MongoDatabase).Collection("dsp")
		cur, err := collection.Find(ctx, bson.D{{}})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer cur.Close(ctx)

		for cur.Next(ctx) {
			//Create a value into which the single document can be decoded
			var elem dsp.DSPCfg
			err := cur.Decode(&elem)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			elements = append(elements, elem)

		}
		if err := cur.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Printf("Found multiple documents: %+v\n", elements)
		res, err := json.Marshal(elements)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(res)
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) addDSP(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dsp dsp.DSPCfg
		fmt.Println(r.Body)
		err := json.NewDecoder(r.Body).Decode(&dsp)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(dsp)
		collection := s.mongo.MongoClient.Database(s.cfg.MongoDatabase).Collection("dsp")

		query := bson.M{"id": bson.M{"$eq": dsp.ID}}

		update := bson.M{"$set": bson.M{
			"id":       dsp.ID,
			"name":     dsp.Name,
			"endpoint": dsp.Endpoint,
			"type":     dsp.Type,
			"qps":      dsp.QPS,
		}}

		opts := options.Update().SetUpsert(true)

		result, err := collection.UpdateOne(ctx, query, update, opts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(result)
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) deleteDSP(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("id")

		fmt.Println(id)
		i, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		collection := s.mongo.MongoClient.Database(s.cfg.MongoDatabase).Collection("dsp")

		result, err := collection.DeleteOne(ctx, bson.M{"id": bson.M{"$eq": i}})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)
		w.WriteHeader(http.StatusOK)
	}
}
