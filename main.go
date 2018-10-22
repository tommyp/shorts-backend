package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mlbright/forecast/v2"
)

func main() {
	r := gin.Default()
	r.GET("/forecast", func(c *gin.Context) {
		lat := c.Query("lat")
		lng := c.Query("lng")

		key := os.Getenv("FORECAST_API_KEY")

		f, err := forecast.Get(key, lat, lng, "now", forecast.UK, forecast.English)
		if err != nil {
			log.Println(err)
		}

		c.JSON(200, f)
	})
	r.Run()
}
