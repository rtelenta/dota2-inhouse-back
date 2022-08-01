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

	player.Active = true

	var newid primitive.ObjectID

	if newid, err = playersDB.Create(newPlayer); err != nil {
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

func PlayersDetails(playerId string) (player domain.Player, code int, err error) {
	if player, err = playersDB.Details(playerId); err != nil {
		code = http.StatusServiceUnavailable
		err = errors.New(http.StatusText(code))
		return
	}

	return
}

func PlayersDelete(playerId string) (deletedPlayer domain.Player, code int, err error) {
	if deletedPlayer, code, err = PlayersDetails(playerId); err != nil {
		return
	}

	if err = playersDB.Delete(playerId); err != nil {
		code = http.StatusServiceUnavailable
		err = errors.New(http.StatusText(code))
		return
	}

	return
}

func PlayersUpdate(player domain.Player) (updatedPlayer domain.Player, code int, err error) {

	player.Active = true

	if updatedPlayer, err = playersDB.Update(player); err != nil {
		code = http.StatusServiceUnavailable
		err = errors.New(http.StatusText(code))
		return
	}

	return
}
