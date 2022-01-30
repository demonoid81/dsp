package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"sync"
	"time"

	"github.com/demonoid81/dsp/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Config["mongo_url"].(string)))
	if err != nil {
		return nil, err
	}
	return client, nil
}




func UpdateReqSetClick(ctx context.Context, key string, waitGroup *sync.WaitGroup, client *mongo.Client) {
	defer waitGroup.Done()
	collection := client.Database(config.Config["mongo_database"].(string)).Collection(config.Config["mongo_collection"].(string))
	filter := bson.M{
		"uuid": bson.M{
			"$eq": key, // check if bool field has value of 'false'
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