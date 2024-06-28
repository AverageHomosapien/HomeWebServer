package main

import (
	"fmt"
	"net/http"
  	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main(){
	fmt.Println("Hello, World!")
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", showIndexPage)

	router.Run()
}

func showIndexPage(c *gin.Context) {

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title": "My Super Special Home Page",
		},
	)
}