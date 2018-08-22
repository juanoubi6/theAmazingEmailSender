package sendgrid

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"theAmazingEmailSender/app/config"
	"theAmazingEmailSender/app/communication/nats/messages"
)

var OwnerEmail string = "contact@theAmazingCodeExaple.com"
var OwnerName string = "The Amazing Code Example"

type EmailSendMessage struct{
	Subject 	string `json:"subject"`
	Message 	string `json:"message"`
	Username	string `json:"user_name"`
	UserEmail	string `json:"user_email"`
}

func SendGenericIndividualEmail(emailData messages.IndividualEmailSendRequest) error {

	from := mail.NewEmail(OwnerName, OwnerEmail)
	subject := emailData.Subject
	to := mail.NewEmail(emailData.UserName, emailData.UserEmail)
	plainTextContent := "The Amazing Code Example"
	htmlContent := emailData.Message

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(config.GetConfig().SENDGRID_KEY_ID)

	response, err := client.Send(message)
	if err != nil {
		println(err.Error())
		return err
	} else {
		println(response.Body)
		return nil
	}

}
