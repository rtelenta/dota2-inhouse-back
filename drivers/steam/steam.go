package steam

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"time"

	"renzotelenta.com/dota2/domain"
)

type Handler struct{}

func NewHandler() (h Handler) {
	h = Handler{}
	return
}

type VanityUrlResource struct {
	Response struct {
		SteamId string `json:"steamid"`
		Success int    `json:"success"`
	} `json:"response"`
}

type PlayerDataResource struct {
	Response struct {
		Players domain.Players `json:"players"`
	} `json:"response"`
}

var key string

func init() {

	if os.Getenv("STEAM_API_KEY") != "" {
		key = os.Getenv("STEAM_API_KEY")
	}

}

func (Handler) GetSteamId(vanityurl string) (steamId string, err error) {
	var isSteamId bool

	isSteamId, _ = regexp.MatchString("^[0-9]{17}$", vanityurl)

	if isSteamId {
		steamId = vanityurl
		return
	}

	statusVanityUrl, bodyVanityUrl, errVanityUrl := exec("GET", fmt.Sprintf("%s%s%s%s", "/ISteamUser/ResolveVanityURL/v0001/?key=", key, "&vanityurl=", vanityurl), nil)

	if errVanityUrl != nil {
		return
	}

	if statusVanityUrl != http.StatusOK {
		err = fmt.Errorf("status code error from steam: %v", statusVanityUrl)
		return
	}

	var vanityUrlResource VanityUrlResource

	json.Unmarshal(bodyVanityUrl, &vanityUrlResource)

	var okResponse = 1

	if vanityUrlResource.Response.Success != okResponse {
		err = fmt.Errorf("error occurred at: %v", time.Now())
		return
	}

	steamId = vanityUrlResource.Response.SteamId

	return
}

func (Handler) GetPlayerData(steamId string) (player domain.Player, err error) {
	statusPlayerData, bodyPlayerData, errPlayerData := exec("GET", fmt.Sprintf("%s%s%s%s", "/ISteamUser/GetPlayerSummaries/v2/?format=json&key=", key, "&steamids=", steamId), nil)

	if errPlayerData != nil {
		return
	}

	if statusPlayerData != http.StatusOK {
		err = fmt.Errorf("status code error from steam: %v", statusPlayerData)
		return
	}

	var playerData PlayerDataResource

	json.Unmarshal(bodyPlayerData, &playerData)

	player = playerData.Response.Players[0]

	return
}
