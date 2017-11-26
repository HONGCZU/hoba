package main

import (
	"flag"
	"fmt"
	"net/http"

	// "github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

var (
	port = flag.Int("port", 8080, "port")
)

func main() {
	flag.Parse()
	router := gin.Default()
	router.LoadHTMLFiles("index.html")
	router.Static("/html", "./html")
	router.Static("/img", "./img")
	router.Static("/css", "./css")
	router.Static("/js", "./js")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	fmt.Println(*port)

	router.Run(fmt.Sprintf(":%d", *port))
}
