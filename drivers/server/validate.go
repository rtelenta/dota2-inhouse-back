package server

import (
	"net/http"

	"renzotelenta.com/dota2/domain"

	"github.com/gin-gonic/gin"
)

func validPlayer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var player domain.Player
		if err := c.ShouldBindJSON(&player); err != nil {
			c.Errors = append(c.Errors, &gin.Error{Err: err})
			code := http.StatusBadRequest
			c.AbortWithStatusJSON(code, gin.H{
				"code":  code,
				"error": err.Error(),
			})
			return
		}

		c.Set("player", player)
		c.Next()
	}
}
