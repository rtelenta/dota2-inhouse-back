package usecases

import (
	"errors"
	"net/http"

	"renzotelenta.com/dota2/domain"
)

func PlayerDetails(vanityurl string) (newPlayer domain.Player, code int, err error) {
	var player domain.Player

	if player, err = steamAPI.GetPlayerData(vanityurl); err != nil {
		code = http.StatusServiceUnavailable
		err = errors.New(http.StatusText(code))
		return
	}

	newPlayer = player

	return
}
