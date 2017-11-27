package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var (
	port = flag.Int("port", 8080, "port")

	// mailgun
	domain       = "mail.17.media"
	apiKey       = "key-346baf7676f148c9ba82d23789a4b0ef"
	publicApiKey = "pubkey-d42e7b0190cc57d61bee2b931e936715"

	// sendgrid
	sendgridApiKey = "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"
)

func main() {
	flag.Parse()
	r := gin.Default()
	r.LoadHTMLFiles("index.html", "html/product.html", "html/news.html", "html/contact.html")
	r.Static("/html", "./html")
	r.Static("/img", "./img")
	r.Static("/css", "./css")
	r.Static("/js", "./js")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.GET("/product", func(c *gin.Context) {
		c.HTML(http.StatusOK, "product.html", gin.H{})
	})
	r.GET("/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.html", gin.H{})
	})
	r.GET("/contact", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contact.html", gin.H{})
	})
	r.POST("/send", send)

	r.Run(fmt.Sprintf(":%d", *port))
}

func send(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	message := c.PostForm("message")

	// mg := mailgun.NewMailgun(domain, apiKey, publicApiKey)
	// msg := mailgun.NewMessage(
	//     "jacky.c@17.media",
	//     "Hoba contact mail",
	//     fmt.Sprintf("name: \n\t%s\nemail: \n\t%s\nphone: \n\t%s\nmessage: \n\t%s\n", name, email, phone, message),
	//     "jacky.c@17.media",
	// )
	// resp, id, err := mg.Send(msg)
	// if err != nil {
	//     log.Fatal(err)
	// }
	// fmt.Printf("ID: %s Resp: %s\n", id, resp)

	from := mail.NewEmail("Jacky Chen", "ps10659@gmail.com")
	subject := "Hoba website's contact mail"
	to := mail.NewEmail("Jacky Chen", "ps10659@gmail.com")
	plainTextContent := fmt.Sprintf("name: \n\t%s\nemail: \n\t%s\nphone: \n\t%s\nmessage: \n\t%s\n", name, email, phone, message)
	msg := mail.NewSingleEmail(from, subject, to, plainTextContent, "")
	client := sendgrid.NewSendClient(os.Getenv(sendgridApiKey))
	response, err := client.Send(msg)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

	fmt.Printf("name: \n\t%s\nemail: \n\t%s\nphone: \n\t%s\nmessage: \n\t%s\n", name, email, phone, message)

	c.Redirect(http.StatusMovedPermanently, "/contact")
	return
}
