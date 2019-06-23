package main

import (
	"github.com/gin-gonic/gin"

	"fetcher"
)

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/forecast.json", func(c *gin.Context) {
		lat := c.Query("lat")
		lng := c.Query("lng")

		new := c.Query("new")

		q := fetcher.Query{
			Latitude:  lat,
			Longitude: lng,
		}

		f := fetcher.GetWeather(q)

		c.JSON(200, f)
	})
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.Abort()
			return
		}
		c.Next()
	}
}
