package database

import (
	"context"
	"log"

	"github.com/maikpro/web_opt_mangadownloader/models"
	"github.com/maikpro/web_opt_mangadownloader/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToMongoDB(ctx context.Context) (*mongo.Client, error) {
	connectionString, err := util.GetEnvString("MONGODB_CONNECTION_STRING")
	if err != nil {
		return nil, err
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return client, nil
}

func getCollection(client *mongo.Client) (*mongo.Collection, error) {
	databaseString, err := util.GetEnvString("MONGODB_DATABASE_STRING")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	collectionString, err := util.GetEnvString("MONGODB_COLLECTION_STRING")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client.Database(databaseString).Collection(collectionString), nil
}

func Save(settings models.Settings) (*models.Settings, error) {
	ctx := context.TODO()
	client, err := connectToMongoDB(ctx)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer client.Disconnect(ctx)
	collection, err := getCollection(client)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	_, err = collection.InsertOne(ctx, settings)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return &settings, nil
}
