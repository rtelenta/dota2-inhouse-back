package mongodb

import (
	"context"
	"time"

	"renzotelenta.com/dota2/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PlayersHandler struct{}

var players *mongo.Collection

func (PlayersHandler) Create(player domain.Player) (newid primitive.ObjectID, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	res, err := players.InsertOne(ctx, player)
	newid = res.InsertedID.(primitive.ObjectID)
	return
}

func (PlayersHandler) Exists(steamId string) (exists bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"steam_id": steamId, "active": true}

	count, err := players.CountDocuments(ctx, filter)

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

	filter := bson.M{"active": true}

	cur, err := players.Find(ctx, filter)

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

func (PlayersHandler) Details(playerId string) (player domain.Player, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var oid primitive.ObjectID

	oid, err = primitive.ObjectIDFromHex(playerId)

	filter := bson.M{"active": true, "_id": oid}

	err = players.FindOne(ctx, filter).Decode(&player)

	return
}

func (PlayersHandler) Delete(playerId string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(playerId)

	if err != nil {
		return
	}

	update := bson.M{
		"$set": bson.M{
			"active": false,
		},
	}

	_, err = players.UpdateByID(ctx, oid, update)

	return
}

func (PlayersHandler) Update(player domain.Player) (updatedPlayer domain.Player, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err != nil {
		return
	}

	filter := bson.M{
		"_id": player.ID,
	}

	update := bson.M{
		"$set": player,
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err = players.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedPlayer)

	return
}
