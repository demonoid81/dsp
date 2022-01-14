package server

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
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
	AdIcon            string            `json:"ad_icon" bson:"ad_icon"`
	AdText            string            `json:"ad_text" bson:"ad_text"`
	Browser           []string          `json:"browser" bson:"browsers"`
	UserID            int               `json:"user_id" bson:"user_id"`
	AdImage           string            `json:"ad_image" bson:"ad_image"`
	AdTitle           string            `json:"ad_title" bson:"ad_title"`
	Category          int               `json:"category" bson:"category"`
	Blacklist         []string          `json:"blacklist" bson:"blacklist"`
	CampaignCountries []CampaignCountry `json:"campaign_country" bson:"campaign_country"`
	Freshness         CampaignFreshness `json:"freshness" bson:"freshness"`
	Whitelist         []string          `json:"whitelist" bson:"whitelist"`
	BlacklistFeed     []int             `json:"blacklist_feed" bson:"blacklist_feed"`
	WhitelistFeed     []int             `json:"whitelist_feed" bson:"whitelist_feed"`
}

var campaigns = Campaign{

	ID:  161,
	OS:  []string{"Windows"},
	URL: "https://poreztaranom.xyz/click.php?key=bxvz5chbvddg7tobrdz9&b={COST}&s={SOURCE_ID}&c={CAMPAIGN_ID}&f={FRESHNESS}&fd={FEED_ID}",
	Type: СampaignType{
		InPage:  false,
		Classic: true,
	},
	AdIcon:   "yq3KNC1otiXWXaJ.png",
	AdText:   "Click here to clean",
	Browser:  []string{},
	UserID:   1261,
	AdImage:  "eaqfaiw6EZDvmFs.png",
	AdTitle:  "System is infected!",
	Category: 1,
	Blacklist: []string{
		"102101neutral__w10_0608_wwcpa",
		"102101terame_w10_3008_ch1_479567",
		"102101sam_w10_2508_multich",
		"102101agent_w10_0507_wwcpa",
		"102101ramos_w10_0907_us_chrome_block",
		"102101ramos_w10_0605_us_chrome",
		"102101ramos_w10_1404_back_new",
		"102101glx_w10_1606_phoenix5",
		"102101medo_w10_0610_multich2",
		"",
	},
	CampaignCountries: []CampaignCountry{
		{Country: "US", CPC: 0.18},
		{Country: "RU", CPC: 0.18},
	},
	Freshness:     CampaignFreshness{Type: "m", Duration: 60},
	Whitelist:     []string{},
	BlacklistFeed: []int{},
	WhitelistFeed: []int{101},
}

//func (s *Server) addСampaign(ctx context.Context) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		collection := s.mongo.MongoClient.Database(s.cfg.MongoDatabase).Collection("campaigns")
//		_, err := collection.InsertOne(ctx, campaigns)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusBadRequest)
//			return
//		}
//		w.WriteHeader(http.StatusOK)
//	}
//}

func (s *Server) addCampaign(ctx context.Context) {
	collection := s.mongo.MongoClient.Database(s.cfg.MongoDatabase).Collection("campaigns")
	_, err := collection.InsertOne(ctx, campaigns)
	if err != nil {
		fmt.Println(err)
		return
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
