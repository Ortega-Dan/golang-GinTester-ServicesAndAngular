package main

import (
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var staticDir string = "../../typescript/ePortfolio/dist/ePortfolio/"

func main() {

	// setting up gin server
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// Serving static content
	r.Use(gin.Recovery(), static.Serve("/", static.LocalFile(staticDir, false)))
	r.NoRoute(func(c *gin.Context) {
		c.File(staticDir + "index.html")
	})

	// Adding endpoints
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// running it
	r.Run()
}
