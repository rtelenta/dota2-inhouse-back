package server

import (
	"net/http"

	"renzotelenta.com/dota2/domain"
	"renzotelenta.com/dota2/usecases"

	"github.com/gin-gonic/gin"
)

func postPlayers(c *gin.Context) {
	player := c.MustGet("player").(domain.Player)
	newPlayer, code, err := usecases.PlayersCreate(player)
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

func getPlayers(c *gin.Context) {
	players, code, err := usecases.PlayersList()
	if err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		c.JSON(code, gin.H{
			"code":  code,
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, players)
	}
}
