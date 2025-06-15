package repository

import (
	"context"
	"log"

	"github.com/maikpro/web_opt_mangadownloader/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ISettingsRepository interface {
	InsertOne(ctx context.Context, settings *models.Settings) (*models.Settings, error)
	FindFirst(ctx context.Context) (*models.Settings, error)
	UpdateOne(ctx context.Context, settings models.Settings) (*models.Settings, error)
}

type SettingsRepository struct {
	collection *mongo.Collection
}

func (settingsRepository *SettingsRepository) InsertOne(ctx context.Context, settings *models.Settings) (*models.Settings, error) {
	newObjectIdString := primitive.NewObjectID().Hex()
	settings.ID = newObjectIdString
	_, err := settingsRepository.collection.InsertOne(ctx, settings)
	if err != nil {
		return nil, err
	}

	return settings, nil
}

func (settingsRepository *SettingsRepository) FindFirst(ctx context.Context) (*models.Settings, error) {
	var settings models.Settings
	options := options.FindOne().SetSort(map[string]int{"_id": -1})
	err := settingsRepository.collection.FindOne(ctx, bson.M{}, options).Decode(&settings)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Fatal(err)
		return nil, err
	}
	return &settings, nil
}
