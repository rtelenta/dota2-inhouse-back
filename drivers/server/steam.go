package server

import (
	"net/http"

	"renzotelenta.com/dota2/usecases"

	"github.com/gin-gonic/gin"
)

func getSteamUser(c *gin.Context) {
	playerUrl := c.Param("playerUrl")
	newPlayer, code, err := usecases.PlayerDetails(playerUrl)
	if err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		c.JSON(code, gin.H{
			"code":  code,
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, newPlayer)
	}
}
