package db

import (
	"context"
	"time"

	"test/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

type Helper struct {
	DB *mongo.Client
}

type Model struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

func ConnectDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opts := options.Client()
	opts.Monitor = otelmongo.NewMonitor()

	client, err := mongo.Connect(ctx, opts.ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Log.Fatal(err)
	}
	log.Log.Println("Connected to MongoDB!")
	return client

}

func (h Helper) GetCollection(dbName string, collectionName string) *mongo.Collection {
	return h.DB.Database(dbName).Collection(collectionName)
}

func (h Helper) CreateCollection(dbName string, collectionName string) {

	docs := []interface{}{
		bson.D{{"id", "1"}, {"title", "Buy  groceries"}},
		bson.D{{"id", "2"}, {"title", "install Aspecto.io"}},
		bson.D{{"id", "3"}, {"title", "Buy dogz.io domain"}},
	}
	h.DB.Database(dbName).Collection(collectionName).InsertMany(context.Background(), docs)
}
