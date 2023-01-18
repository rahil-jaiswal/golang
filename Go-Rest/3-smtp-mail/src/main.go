package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/mail"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
)

func composeMail(name string, email string, contact string, query string) (string, int, error) {

	credFile, err := os.Open("credentials.json")

	if err != nil {
		return "Passkey Not Found", 500, err
	}
	defer credFile.Close()

	credBytes, err := ioutil.ReadAll(credFile)
	var credJson map[string]interface{}
	json.Unmarshal([]byte(credBytes), &credJson)

	password, err := base64.StdEncoding.DecodeString(credJson["passkey"].(string))
	if err != nil {
		return "Password Decoding Failed", 500, err
	}

	// smtp server configuration.
	configFile, err := os.Open("config.json")
	if err != nil {
		return "Config File Not Found", 500, err
	}
	defer configFile.Close()

	configStrings, err := ioutil.ReadAll(configFile)
	var configJson map[string]interface{}
	json.Unmarshal([]byte(configStrings), &configJson)
	smtpHost := configJson["smtpHost"].(string)
	smtpPort := configJson["smtpPort"].(string)
	senderEmail := credJson["email"].(string)
	from := mail.Address{Name: configJson["senderName"].(string), Address: senderEmail}
	to := mail.Address{Name: configJson["receiverName"].(string), Address: configJson["receiverEmail"].(string)}
	body := "New Enquiry. Following are enquirer details. \n \t Name\t: \t" + name + "\n\t Email\t:\t" + email + "\n \t Contact \t:\t" + contact + "\n\t Query\t:\t" + query + "\nPlease ask responsible representative of CoffeeBeans to connect with them.\n"

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = "New Enquiry by " + name + " (" + email + ")"
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	auth := smtp.PlainAuth("", senderEmail, string(password[:]), smtpHost)

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from.Address, []string{to.Address}, []byte(message))

	if err != nil {
		return "Error Sending Mail", 503, err
	}
	return "Email Sent Successfully!", 200, nil
}

func sendEnquiry(c *gin.Context) {
	body := c.Request.Body
	reuqestJson, err := ioutil.ReadAll(body)
	var requestMap map[string]interface{}

	json.Unmarshal([]byte(reuqestJson), &requestMap)
	enquirer_name := requestMap["name"].(string)
	enquirer_email := requestMap["email"].(string)
	enquirer_contact := requestMap["contact"].(string)
	enquirer_query := requestMap["query"].(string)

	responseMessage, responseCode, err := composeMail(enquirer_name, enquirer_email, enquirer_contact, enquirer_query)
	fmt.Print(err)
	if err != nil {
		fmt.Println(err)
		c.JSON(501, gin.H{
			"message": "501 - Server-Side Problem",
		})
		return
	}
	c.JSON(responseCode, gin.H{
		"message": responseMessage,
	})
}

func main() {
	fmt.Println("Hello world!")
	r := gin.Default()
	r.POST("/sendEnquiry", sendEnquiry)
	r.Run(":8001")
}
