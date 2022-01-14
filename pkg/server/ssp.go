package server

//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"github.com/demonoid81/dsp/auction/dsp"
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/mongo"
//	"net/http"
//)
//
//func (s *Server) loadSSP(ctx context.Context) error {
//	var ssp []dsp.SSP
//
//	collection := s.mongo.MongoClient.Database(s.cfg.MongoDatabase).Collection("ssp")
//	cur, err := collection.Find(ctx, bson.D{{}})
//	if err != nil {
//		return err
//	}
//	defer cur.Close(ctx)
//
//	for cur.Next(ctx) {
//		//Create a value into which the single document can be decoded
//		var elem dsp.SSP
//		err := cur.Decode(&elem)
//		if err != nil {
//			return err
//		}
//
//		ssp = append(ssp, elem)
//
//	}
//	if err := cur.Err(); err != nil {
//		return err
//	}
//	fmt.Printf("Found multiple documents: %+v\n", ssp)
//	s.arraySSP= ssp
//	return nil
//}
//
//func (s *Server) getSSP(ctx context.Context) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var ssp []dsp.SSP
//
//		collection := s.mongo.MongoClient.Database(s.cfg.MongoDatabase).Collection("ssp")
//		cur, err := collection.Find(ctx, bson.D{{}})
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		defer cur.Close(ctx)
//
//		for cur.Next(ctx) {
//			//Create a value into which the single document can be decoded
//			var elem dsp.SSP
//			err := cur.Decode(&elem)
//			if err != nil {
//				http.Error(w, err.Error(), http.StatusInternalServerError)
//				return
//			}
//
//			ssp = append(ssp, elem)
//
//		}
//		if err := cur.Err(); err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		fmt.Printf("Found multiple documents: %+v\n", ssp)
//		res, err := json.Marshal(ssp)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		w.Write(res)
//		w.WriteHeader(http.StatusOK)
//	}
//}
//
//func (s *Server) addSSP(ctx context.Context, client *mongo.Client) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var ssp dsp.SSP
//		err := json.NewDecoder(r.Body).Decode(&ssp)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusBadRequest)
//			return
//		}
//		fmt.Println(ssp)
//		collection := s.mongo.MongoClient.Database(s.cfg.MongoDatabase).Collection("ssp")
//		_, err = collection.InsertOne(ctx, ssp)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusBadRequest)
//			return
//		}
//		w.WriteHeader(http.StatusOK)
//	}
//}
//
//func reload(ctx context.Context, client *mongo.Client) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		w.WriteHeader(http.StatusOK)
//	}
//}
