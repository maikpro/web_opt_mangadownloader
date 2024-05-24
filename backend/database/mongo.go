package database

import (
	"context"
	"log"

	"github.com/maikpro/web_opt_mangadownloader/models"
	"github.com/maikpro/web_opt_mangadownloader/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func getCollection(client *mongo.Client, selectedCollection string) (*mongo.Collection, error) {
	databaseString, err := util.GetEnvString("MONGODB_DATABASE_STRING")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	collectionString, err := util.GetEnvString(selectedCollection)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client.Database(databaseString).Collection(collectionString), nil
}

func SaveSettings(settings models.Settings) (*models.Settings, error) {
	ctx := context.TODO()
	client, err := connectToMongoDB(ctx)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer client.Disconnect(ctx)
	collection, err := getCollection(client, "MONGODB_COLLECTION_SETTINGS_STRING")
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	newObjectIdString := primitive.NewObjectID().Hex()
	log.Println(newObjectIdString)

	settings.ID = newObjectIdString
	_, err = collection.InsertOne(ctx, settings)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return &settings, nil
}

func GetSettings() (*models.Settings, error) {
	ctx := context.TODO()
	client, err := connectToMongoDB(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer client.Disconnect(ctx)

	options := options.FindOne().SetSort(map[string]int{"_id": -1})

	collection, err := getCollection(client, "MONGODB_COLLECTION_SETTINGS_STRING")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var result models.Settings
	err = collection.FindOne(context.Background(), bson.M{}, options).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Fatal(err)
		return nil, err
	}
	return &result, nil
}

func UpdateSettings(settings models.Settings) (*models.Settings, error) {
	ctx := context.TODO()
	client, err := connectToMongoDB(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer client.Disconnect(ctx)

	filter := bson.M{"_id": settings.ID}
	update := bson.M{
		"$set": settings,
	}
	opts := options.Update().SetUpsert(true)

	collection, err := getCollection(client, "MONGODB_COLLECTION_SETTINGS_STRING")
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	_, err = collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &settings, nil
}
