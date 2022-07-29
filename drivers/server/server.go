package server

import (
	"net/http"

	"renzotelenta.com/dota2/drivers/stats"

	"github.com/gin-gonic/gin"
)

func NewEngine() (engine *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)
	engine = gin.New()
	return
}

func Setup(engine *gin.Engine) {
	engine.Use(gin.Recovery())
	setRoutes(engine)
}

func setRoutes(engine *gin.Engine) {
	engine.GET("/", status)
	engine.GET("/health", health)

	api := engine.Group("/api")
	api.POST("/players", validPlayer(), postPlayers)
	api.GET("/players", getPlayers)

	steam := api.Group("/steam")
	steam.GET("/players/:playerUrl", getSteamUser)
}

func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "up"})
}

func health(c *gin.Context) {
	status := stats.GetOsStats()
	c.JSON(http.StatusOK, status)
}
