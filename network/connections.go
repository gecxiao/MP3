package network

import (
	"../application"
	"fmt"
	"log"
	"math"
	"net"
)


func Initialize(processes []application.Process) []application.Process{
	//initialize tcp connection for all the processes
	for i :=0 ; i<len(processes); i++{
		ln, err := net.Listen("tcp", ":"+ processes[i].Port)
		if err != nil{
			log.Panic(err)
		}
		processes[i].Ln = ln
	}
	for i:=0 ; i<len(processes); i++{
		for j := 0; j<len(processes); j++{
			c, err := net.Dial("tcp", processes[j].Ip+":"+processes[j].Port)
			if err != nil{
				log.Panic(err)
			}
			processes[i].Conns = append(processes[i].Conns, c)
		}
	}
	return processes
}

//func handleConnection(c net.Conn, messages chan application.Message) {
//	for {
//		decoder := gob.NewDecoder(c)
//		mes := new(application.Message)
//		_ = decoder.Decode(mes)
//
//		messages<-*mes
//		}
//}

func ProcessMsg(ln net.Listener, messages chan application.Message) {
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go UnicastReceive(c, messages)
	}
}

func ApproximateConsensus(messages chan application.Message, process application.Process, fNum int, totalNum int,
	maxDelay int, minDelay int, control chan bool) {
	//implement the approximate consensus algorithm
	r := 1
	sum := 0.
	msgNum := 0
	valR1 := process.InitialV
	firstMsg := new(application.Message)
	firstMsg.M = process.InitialV
	firstMsg.S = process.Id
	firstMsg.R = 1
	MulticastSend(process.Conns, *firstMsg, maxDelay, minDelay)
	var tooEarly []application.Message
	for {
		message := <- messages
		if message.R < r {
			continue
		} else if message.R == r {
			msgNum++
			sum += message.M
		} else if message.R > r {
			tooEarly = append(tooEarly, message)
		}

		if msgNum == totalNum-fNum {
			mes := new(application.Message)
			valR2 := sum / float64(totalNum-fNum)
			mes.M = valR2
			mes.S = process.Id
			mes.R = r+1
			MulticastSend(process.Conns, *mes, maxDelay, minDelay)
			for _, message := range tooEarly {
				messages <- message
			}
			_, err := fmt.Printf("Process %s has value %.3f at round %d \n", process.Id, mes.M, r)
			if err!=nil{
				panic(err)
			}
			if math.Abs(valR1 - valR2) < 0.001 {
				control <- true
				break
			}
			sum = 0
			msgNum = 0
			r++
			valR1 = valR2
		}
	}
}