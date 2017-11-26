package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mailgun/mailgun-go.v1"

	"github.com/gin-gonic/gin"
)

var (
	port = flag.Int("port", 8080, "port")

	// mailgun
	domain       = "mail.17.media"
	apiKey       = "key-346baf7676f148c9ba82d23789a4b0ef"
	publicApiKey = "pubkey-d42e7b0190cc57d61bee2b931e936715"
)

func main() {
	flag.Parse()
	r := gin.Default()
	r.LoadHTMLFiles("index.html")
	r.Static("/html", "./html")
	r.Static("/img", "./img")
	r.Static("/css", "./css")
	r.Static("/js", "./js")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/html/contact.html", gin.H{})
	})
	r.POST("/send", send)

	r.Run(fmt.Sprintf(":%d", *port))
}

func send(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	message := c.PostForm("message")
	c.Redirect(http.StatusMovedPermanently, "html/contact.html")
	fmt.Printf("email: %s\nname: %s\nphone: %s\nmessage: %s", email, name, phone, message)

	mg := mailgun.NewMailgun(domain, apiKey, publicApiKey)
	msg := mailgun.NewMessage(
		"ps10659@gmail.com",
		"Hoba contact mail",
		fmt.Sprintf("email: %s\nname: %s\nphone: %s\nmessage: %s", email, name, phone, message),
		"ps10659@gmail.com",
	)
	resp, id, err := mg.Send(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %s Resp: %s\n", id, resp)

	return
}
