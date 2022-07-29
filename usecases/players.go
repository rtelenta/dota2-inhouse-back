package usecases

import (
	"errors"
	"fmt"
	"net/http"

	"renzotelenta.com/dota2/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PlayersCreate(player domain.Player) (newPlayer domain.Player, code int, err error) {

	exists := playersDB.Exists(player.SteamId)

	if exists {
		code = http.StatusConflict
		err = fmt.Errorf("%s - %s is already in use", player.PersonaName, player.ProfileUrl)
		return
	}

	var newid primitive.ObjectID
	if newid, err = playersDB.Create(player); err != nil {
		code = http.StatusServiceUnavailable
		err = errors.New(http.StatusText(code))
		return
	}
	newPlayer = player
	newPlayer.ID = &newid

	return
}

func PlayersList() (players domain.Players, code int, err error) {
	if players, err = playersDB.List(); err != nil {
		code = http.StatusServiceUnavailable
		err = errors.New(http.StatusText(code))
		return
	}

	return
}
