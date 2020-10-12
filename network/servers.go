package network

import (
	"../application"
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

func initialize(processes []application.Process){
	for i :=0 ; i<len(processes); i++{
		ln, err := net.Listen("tcp", ":"+ processes[i].Port)
		if err != nil{
			log.Panic(err)
		}
		processes[i].Ln = ln
	}
	for j := 0; j<len(processes); j++{
		c, err := net.Dial("tcp", processes[j].Ip+":"+processes[j].Port)
		if err != nil{
			log.Panic(err)
		}
		processes[j].Conns = append(processes[j].Conns, c)
	}
}
func handleConnection(c net.Conn, messages chan application.Message) {
	for {
		decoder := gob.NewDecoder(c)
		mes := new(application.Message)
		_ = decoder.Decode(mes)
		messages<-*mes
		}
}

func Server(server application.Process, n int, messages chan application.Message) {
	//input: the network# and the # of connections it will receive
	//listen to the client and decode the application, then send via channel
	//simulate the delay here.
	var counter = 0
	ln, err := net.Listen("tcp", server.Ip+":"+server.Port) //creates server
	if err != nil {
		fmt.Println(err)
	}
	defer ln.Close()
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		if counter == n {
			err = c.Close()
			if err != nil {
				fmt.Println(err)
			}
			return
		}
		go handleConnection(c, messages)
	}
}
