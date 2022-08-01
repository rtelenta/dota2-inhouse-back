package server

import (
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func getPlayersById(c *gin.Context) {
	id := c.Param("playerId")

	player, code, err := usecases.PlayersDetails(id)

	if err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		c.JSON(code, gin.H{
			"code":  code,
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, player)
	}
}

func deletePlayers(c *gin.Context) {
	id := c.Param("playerId")

	deletedPlayer, code, err := usecases.PlayersDelete(id)

	if err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		c.JSON(code, gin.H{
			"code":  code,
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, deletedPlayer)
	}
}

func putPlayers(c *gin.Context) {
	id := c.Param("playerId")
	player := c.MustGet("player").(domain.Player)
	oid, err := primitive.ObjectIDFromHex(id)

	player.ID = &oid

	updatedPlayer, code, err := usecases.PlayersUpdate(player)

	if err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		c.JSON(code, gin.H{
			"code":  code,
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, updatedPlayer)
	}
}
