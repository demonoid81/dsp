package app

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

type Mongo struct {
	MongoClient *mongo.Client
	cfg         *Config
}

func (app *Env) NewMongoClient() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(app.Cfg.MongoURL))
	if err != nil {
		return err
	}
	app.Mongo = &Mongo{
		MongoClient: client,
		cfg:         app.Cfg,
	}
	return nil
}

func (client *Mongo) AddDSPRequest(ctx context.Context, data interface{}, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	collection := client.MongoClient.Database(client.cfg.MongoDatabase).Collection("requests")
	result, err := collection.InsertOne(ctx, data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func (client *Mongo) UpdateDSPRequestSetClick(ctx context.Context, key string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	collection := client.MongoClient.Database(client.cfg.MongoDatabase).Collection("requests")
	filter := bson.M{
		"uuid": bson.M{
			"$eq": key,
		},
	}
	update := bson.M{
		"$set": bson.M{
			"click": true,
		},
	}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
