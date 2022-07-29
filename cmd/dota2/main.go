package main

import (
	_ "github.com/joho/godotenv/autoload"
	"renzotelenta.com/dota2/drivers/mongodb"
	"renzotelenta.com/dota2/drivers/server"
	"renzotelenta.com/dota2/drivers/steam"
	"renzotelenta.com/dota2/usecases"
)

func main() {
	mongoAccess := mongodb.NewHandler()
	steamApi := steam.NewHandler()

	usecases.SetBridges(
		mongoAccess.Players,
		steamApi,
	)

	engine := server.NewEngine()
	server.Setup(engine)
	engine.Run(":3006")

}
