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
	time.Sleep(time.Duration(d) * time.Millisecond)
	encoder := gob.NewEncoder(c)
	err := encoder.Encode(m)
	if err != nil {
		panic(err)
	}
}

func MulticastSend(conns []net.Conn, m application.Message, maxDelay int, minDelay int){
	for i :=0 ; i< len(conns); i++{
		UnicastSend(conns[i], m,maxDelay, minDelay)
	}
}

func UnicastReceive(c net.Conn, messages chan application.Message) {
	for{
		decoder := gob.NewDecoder(c)
		mes := new(application.Message)
		err := decoder.Decode(mes)
		if err != nil {
			panic(err)
		}
		messages<-*mes
	}
}