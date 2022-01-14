package server
//
//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"github.com/demonoid81/dsp/auction/dsp"
//	"github.com/demonoid81/dsp/config"
//	"go.mongodb.org/mongo-driver/bson"
//	"net/http"
//)
//
//
//
//func (app *server.app) getDSP(ctx context.Context) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var elements []dsp.DSPCfg
//
//		collection := app.mongoClient.Database(config.Config["mongo_database"].(string)).Collection("dsp")
//		cur, err := collection.Find(ctx, bson.D{{}})
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		defer cur.Close(ctx)
//
//		for cur.Next(ctx) {
//			//Create a value into which the single document can be decoded
//			var elem dsp.DSPCfg
//			err := cur.Decode(&elem)
//			if err != nil {
//				http.Error(w, err.Error(), http.StatusInternalServerError)
//				return
//			}
//
//			elements = append(elements, elem)
//
//		}
//		if err := cur.Err(); err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		fmt.Printf("Found multiple documents: %+v\n", elements)
//		res, err := json.Marshal(elements)
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
//func (app *server.app) addDSP(ctx context.Context) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var dsp dsp.DSPCfg
//		err := json.NewDecoder(r.Body).Decode(&dsp)
//		if err != nil {
//			fmt.Println(err)
//			http.Error(w, err.Error(), http.StatusBadRequest)
//			return
//		}
//		fmt.Println(dsp)
//		collection := app.mongoClient.Database(config.Config["mongo_database"].(string)).Collection("dsp")
//		result, err := collection.InsertOne(ctx, dsp)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusBadRequest)
//			return
//		}
//		fmt.Println(result)
//		w.WriteHeader(http.StatusOK)
//	}
//}
//
//func (app *server.app) updateDSP(ctx context.Context) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var dsp dsp.DSPCfg
//		err := json.NewDecoder(r.Body).Decode(&dsp)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusBadRequest)
//			return
//		}
//		fmt.Println(dsp)
//		collection := app.mongoClient.Database(config.Config["mongo_database"].(string)).Collection("dsp")
//
//		result, err := collection.InsertOne(ctx, dsp)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusBadRequest)
//			return
//		}
//		fmt.Println(result)
//		w.WriteHeader(http.StatusOK)
//	}
//}
