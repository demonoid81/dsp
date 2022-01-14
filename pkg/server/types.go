package server

type campany struct {
	UID float64 `json:"uid"`
	Cur string  `json:"cur"`
	Cpr float64 `json:"cpr"`
	Cid string  `json:"cid"`
	Atl string  `json:"atl"`
	Atx string  `json:"atx"`
	Aic string  `json:"aic"`
	Aig string  `json:"aig"`
	Ccr string  `json:"ccr"`
}

type LinkData struct {
	Key    string  `json:"uuid" bson:"uuid"`
	Link   string  `json:"link" bson:"-"`
	Cpc    float64 `json:"cpc" bson:"cpc"`
	Uid    float64 `json:"uid" bson:"uid"`
	Cid    string  `json:"cid" bson:"cid"`
	Cou    string  `json:"cou" bson:"cou"`
	Bro    string  `json:"bro" bson:"bro"`
	Os     string  `json:"os" bson:"os"`
	Sid    string  `json:"sid" bson:"sid"`
	Date   string  `json:"date" bson:"date"`
	Fresh  string  `json:"fresh" bson:"fresh"`
	FeedId string  `json:"feed_id" bson:"feed_id"`
	Click  bool    `json:"-" bson:"click"`
}