package common

import (
	"github.com/nats-io/go-nats"

	"theAmazingEmailSender/app/config"
)

var natsConnection *nats.Conn

func ConnectToNats(){

	nc, err := nats.Connect(config.GetConfig().NATS_URL)
	if err != nil{
		panic(err)
	}

	natsConnection = nc

}

func GetNatsEncodedConnection() *nats.EncodedConn{

	c, err := nats.NewEncodedConn(natsConnection, nats.JSON_ENCODER)
	if err != nil{
		panic(err)
	}

	return c

}

func GetNatsConnection() *nats.Conn{
	return natsConnection
}