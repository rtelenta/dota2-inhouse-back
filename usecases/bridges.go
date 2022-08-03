package usecases

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"renzotelenta.com/dota2/domain"
)

var playersDB playersDBbridge
var steamAPI steamAPIbridge

type playersDBbridge interface {
	Create(player domain.Player) (newid primitive.ObjectID, err error)
	Exists(steamId string) (exists bool)
	List() (playersList domain.Players, err error)
	Delete(playerId string) (err error)
	Update(player domain.Player) (updatedPlayer domain.Player, err error)
	Details(playerId string) (player domain.Player, err error)
}

type steamAPIbridge interface {
	GetPlayerData(steamId string) (player domain.Player, err error)
	GetSteamId(vanityurl string) (steamId string, err error)
}

func SetBridges(
	players playersDBbridge,
	steam steamAPIbridge,
) {
	playersDB = players
	steamAPI = steam
}
