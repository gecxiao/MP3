package network

import (
	"../application"
	"encoding/gob"
	"math/rand"
	"net"
	"time"
)

func UnicastSend(c net.Conn, m application.Message, maxDelay int, minDelay int) {
	//Sends message to the destination process.
	d := rand.Intn(maxDelay-minDelay)+minDelay
	time.Sleep(time.Duration(d))
	encoder := gob.NewEncoder(c)
	msg := application.Message{
		S: m.S,
		R: m.R,
		M: m.M,
	}
	_ = encoder.Encode(msg)
	return
}
