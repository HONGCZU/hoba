package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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

	// write email into local log file
	f, err := os.OpenFile("email.log", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	location, _ := time.LoadLocation("Asia/Taipei")
	now := time.Now().In(location).Format("Mon Jan 02 15:04:05 -0700 2006")
	content := fmt.Sprintf("time: %s\nname: %s\nemail: %s\nphone: %s\nmessage: %s\n\n", now, name, email, phone, message)
	fmt.Println(content)

	if _, err = f.WriteString(content); err != nil {
		panic(err)
	}

	// send email by sendgrid
	from := mail.NewEmail("AnAn", "ps10659@gmail.com")
	subject := "Hoba website's mail from " + name + "(" + email + ")"
	to := mail.NewEmail("HongCzu", "sales.hongczu@gmail.com")
	plainTextContent := content
	htmlContent := "<p>time: " + now + "<br>name: " + name + "<br>email: " + email + "<br>phone: " + phone + "<br>message: " + message + "<br></p>"
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

	c.Redirect(http.StatusMovedPermanently, "/contact")
	return
}
