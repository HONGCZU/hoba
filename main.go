package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var (
	port = flag.Int("port", 8080, "port")
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

	if name == "" && email == "" && phone == "" && message == "" {
		c.Redirect(http.StatusMovedPermanently, "/contact")
		return
	}

	from := mail.NewEmail("AnAn", "ps10659@gmail.com")
	subject := "Hoba website's mail from " + name + "(" + email + ")"
	to := mail.NewEmail("HongCzu", "sales.hongczu@gmail.com")
	plainTextContent := fmt.Sprintf("name: \n\t%s\nemail: \n\t%s\nphone: \n\t%s\nmessage: \n\t%s\n", name, email, phone, message)
	htmlContent := "<p>name: " + name + "<br>email: " + email + "<br>phone: " + phone + "<br>message: " + message + "<br></p>"
	msg := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
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
