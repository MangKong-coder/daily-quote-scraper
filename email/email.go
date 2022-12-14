package email

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func EmailQuotes(imageBody string) {
	// Initialize variables
	email := os.Getenv("EMAIL")
	name := os.Getenv("NAME")
	from := mail.NewEmail(name, email)
	subject := "Daily Quote " + time.Now().Local().Format("January 02, 2006")
	to := mail.NewEmail(name, email)

	// Compose email message
	message := mail.NewSingleEmail(from, subject, to, "something", imageBody)

	// Creates new sendgrid client with API Key\
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	// Sends message
	response, err := client.Send(message)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
