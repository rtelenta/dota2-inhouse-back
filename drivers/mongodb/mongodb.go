package mongodb

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Handler struct {
	Matches *MatchesHandler
	Players *PlayersHandler
}

func NewHandler() (h Handler) {
	connect()
	h = Handler{
		Matches: &MatchesHandler{},
		Players: &PlayersHandler{},
	}
	return
}

func connect() {
	opt := options.Client().ApplyURI(os.Getenv("MONGO_URL")).SetMaxPoolSize(4)
	client, _ := mongo.NewClient(opt)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client.Connect(ctx)
	client.Ping(ctx, readpref.Primary())

	matches = client.Database("dota2").Collection("matches")
	players = client.Database("dota2").Collection("players")

}
