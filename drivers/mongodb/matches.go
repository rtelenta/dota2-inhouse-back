package mongodb

import (
	"context"
	"time"

	"renzotelenta.com/dota2/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MatchesHandler struct{}

var matches *mongo.Collection

func (MatchesHandler) Create(c domain.Match) (newid primitive.ObjectID, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	res, err := matches.InsertOne(ctx, c)
	newid = res.InsertedID.(primitive.ObjectID)
	return
}
