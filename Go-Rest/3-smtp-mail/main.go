package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"strings"
)

func encodeRFC2047(String string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}

func main() {
	// Set up authentication information.

	smtpServer := "smtp.gmail.com"
	auth := smtp.PlainAuth(
		"",
		"rahilravijaiswal@gmail.com",
		"gathwtcnswdqqkcg",
		smtpServer,
	)

	from := mail.Address{Name: "CoffeeBeans Website", Address: "rahilravijaiswal@gmail.com"}
	to := mail.Address{Name: "CoffeeBeans Enquiries", Address: "rahil@coffeebeans.io"}
	title := "Test Mail"

	body := "If Received - Testing Completed 2"

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = title
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		smtpServer+":587",
		auth,
		from.Address,
		[]string{to.Address},
		[]byte(message),
		//[]byte("This is the email body."),
	)
	if err != nil {
		log.Fatal(err)
	}

}
