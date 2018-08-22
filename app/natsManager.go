package app

import (
	"theAmazingEmailSender/app/common"
	"strconv"
	"theAmazingEmailSender/app/config"
	"github.com/nats-io/go-nats"
	"theAmazingEmailSender/app/helpers/sendgrid"
	"encoding/json"
	"theAmazingEmailSender/app/communication/nats/messages"
	log "github.com/sirupsen/logrus"
	"fmt"
)

var workerAmount, _ = strconv.Atoi(config.GetConfig().WORKER_AMOUNT)

func ListenToEmailQueue(){

	var natsConn = common.GetNatsConnection()

	for i := 0; i < workerAmount; i++ {
		go checkEmails(natsConn)
	}

	log.Printf("Waiting for email messages")

}

func checkEmails(conn *nats.Conn){
	conn.QueueSubscribe("emailSender:sendIndividualEmail", "sendIndividualEmail_workers", func(newMessage *nats.Msg) {
		log.Printf("Recibi")
		var individualEmailSendResp messages.IndividualEmailSendResponse
		var individualEmailSendReq messages.IndividualEmailSendRequest
		if err := json.Unmarshal(newMessage.Data, &individualEmailSendReq); err != nil {
			log.WithFields(log.Fields{
				"place": "decoding message body",
			}).Info(err.Error())
			SendErrorResponse(err,individualEmailSendResp,newMessage,conn)
			return
		}

		if err := sendgrid.SendGenericIndividualEmail(individualEmailSendReq); err != nil {
			log.WithFields(log.Fields{
				"place": "email send",
			}).Info(err.Error())
			SendErrorResponse(err,individualEmailSendResp,newMessage,conn)
			return
		}

		messageBody, err := individualEmailSendResp.GetMessageBytes()
		if err != nil{
			log.WithFields(log.Fields{
				"place": "encoding message body",
			}).Info(err.Error())
			SendErrorResponse(err,individualEmailSendResp,newMessage,conn)
			return
		}

		conn.Publish(newMessage.Reply,messageBody)
	})

	//Used to tell nats server that there are already subscribers to this topic
	err := conn.Flush()
	if err != nil {
		fmt.Println("Error flushing nats connection: " + err.Error())
	}
}


func SendErrorResponse(err error,individualEmailSendResp messages.IndividualEmailSendResponse,message *nats.Msg,conn *nats.Conn){

	individualEmailSendResp.Error = err.Error()

	messageBody, _ := individualEmailSendResp.GetMessageBytes()

	conn.Publish(message.Reply,messageBody)

}