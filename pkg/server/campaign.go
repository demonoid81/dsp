package server

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type СampaignType struct {
	InPage  bool `json:"inpage" bson:"inpage"`
	Classic bool `json:"classic" bson:"classic"`
}

type CampaignCountry struct {
	Country string  `json:"country"`
	CPC     float64 `json:"cpc"`
}

type CampaignFreshness struct {
	Type     string `json:"type" bson:"type"`
	Duration int    `json:"duration" bson:"duration"`
}

type Campaign struct {
	ID                int               `json:"id" bson:"id"`
	OS                []string          `json:"os" bson:"os"`
	URL               string            `json:"url" bson:"url"`
	Type              СampaignType      `json:"type" bson:"type"`
	Icon              string            `json:"icon" bson:"icon"`
	Text              string            `json:"text" bson:"text"`
	Browser           []string          `json:"browser" bson:"browsers"`
	UserID            int               `json:"user_id" bson:"user_id"`
	Image             string            `json:"image" bson:"image"`
	Title             string            `json:"title" bson:"title"`
	Category          int               `json:"category" bson:"category"`
	CampaignCountries []CampaignCountry `json:"campaign_country" bson:"campaign_country"`
	Freshness         CampaignFreshness `json:"freshness" bson:"freshness"`
	Whitelist         []string          `json:"whitelist" bson:"whitelist"`
	Blacklist         []string          `json:"blacklist" bson:"blacklist"`
	BlacklistFeed     []int             `json:"blacklist_feed" bson:"blacklist_feed"`
	WhitelistFeed     []int             `json:"whitelist_feed" bson:"whitelist_feed"`
}

func (s *Server) addСampaign(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var campaign Campaign
		err := json.NewDecoder(r.Body).Decode(&campaign)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		collection := s.mongo.MongoClient.Database(s.cfg.MongoDatabase).Collection("campaigns")
		_, err = collection.InsertOne(ctx, campaign)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) getCampaign(ctx context.Context) {
	collection := s.mongo.MongoClient.Database(s.cfg.MongoDatabase).Collection("campaigns")
	filter := bson.M{
		"campaign_country.country": bson.M{"$eq": "US"},
		"$and": []bson.M{
			{"$or": []bson.M{
				{"os": bson.M{"$eq": "Windows"}},
				{"$and": []bson.M{
					{"os": bson.M{"$exists": true}},
					{"os": bson.M{"$size": 0}},
				}},
			}},
			{"$or": []bson.M{
				{"browsers": bson.M{"$eq": "Edge"}},
				{"$and": []bson.M{
					{"browsers": bson.M{"$exists": true}},
					{"browsers": bson.M{"$size": 0}},
				}},
			}},
		},
	}
	var campaigns []Campaign
	cur, err := collection.Find(ctx, filter)

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		//Create a value into which the single document can be decoded
		var elem Campaign
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println(err)
			return
		}

		campaigns = append(campaigns, elem)

	}
	if err := cur.Err(); err != nil {
		fmt.Println(err)
		return
	}
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(campaigns)
}
