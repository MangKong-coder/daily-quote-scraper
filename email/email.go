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
	email := os.Getenv("EMAIL")
	name := os.Getenv("NAME")
	from := mail.NewEmail(name, email)
	subject := "Daily Quote " + time.Now().Local().Format("January 02, 2006")
	to := mail.NewEmail(name, email)
	message := mail.NewSingleEmail(from, subject, to, "something", imageBody)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}