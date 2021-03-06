package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// don't forget to add the last "/" at the end of the path
var staticDir string = "../../../typescript/ePortfolio/dist/ePortfolio/"

func main() {

	// setting up gin server
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// Using Gin Recovery AND Serving static content
	r.Use(gin.Recovery(), static.Serve("/", static.LocalFile(staticDir, false)))
	// Enabling angular router
	r.NoRoute(func(c *gin.Context) {
		c.File(staticDir + "index.html")
	})

	// Adding endpoints
	r.GET("/ping", func(c *gin.Context) {
		// panic("wrong")
		fmt.Println("ping get request")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// running it
	// running with ssl
	// err := r.RunTLS(":8023", "certificate.crt", "server.key")
	err := r.Run(":8023")
	if err != nil {
		panic(err)
	}
}
