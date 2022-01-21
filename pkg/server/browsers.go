package server

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)


func (s *Server) getBrowsers(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var countries []country

		collection := s.mongo.MongoClient.Database(s.cfg.MongoDatabase).Collection("countries")
		cur, err := collection.Find(ctx, bson.D{{}})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer cur.Close(ctx)

		for cur.Next(ctx) {
			//Create a value into which the single document can be decoded
			var elem country
			err := cur.Decode(&elem)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			countries = append(countries, elem)

		}
		if err := cur.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res, err := json.Marshal(countries)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(res)
		w.WriteHeader(http.StatusOK)
	}
}
