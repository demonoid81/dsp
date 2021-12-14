package mongodb

import (
	"context"
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
