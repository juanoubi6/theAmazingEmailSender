package main

import (
	"theAmazingEmailSender/app"
	"theAmazingEmailSender/app/common"
)

func main() {
	common.ConnectToNats()
	keep := make(chan bool)
	go app.ListenToEmailQueue()
	<-keep
}
