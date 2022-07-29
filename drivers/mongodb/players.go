package mongodb

import (
	"context"
	"time"

	"renzotelenta.com/dota2/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlayersHandler struct{}

var players *mongo.Collection

func (PlayersHandler) Create(c domain.Player) (newid primitive.ObjectID, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	res, err := players.InsertOne(ctx, c)
	newid = res.InsertedID.(primitive.ObjectID)
	return
}

func (PlayersHandler) Exists(steamId string) (exists bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	count, err := players.CountDocuments(ctx, bson.M{"steam_id": steamId})

	if err == nil && count >= 1 {
		exists = true
		return
	}

	exists = false
	return
}

func (PlayersHandler) List() (playersList domain.Players, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := players.Find(ctx, bson.D{})

	if err != nil {
		return
	}
	defer cur.Close(ctx)

	if err = cur.All(ctx, &playersList); err != nil {
		return
	}

	err = cur.Err()

	if playersList == nil {
		playersList = domain.Players{}
	}

	return
}
