package main

/*
In this simulation
network 1 always wait a connection from network 2
network 3 always wait two connections: from network 2 and 4
The configuration file set max delay to be 3000(ms), min delay to be 1000(ms).
*/
import (
	"./application"
	"./network"
	"./utils"
	"log"
	"time"
)

func main() {
	//get parameters
	minDelay, maxDelay, fNum := utils.FetchDelay()
	counter:=0
	var ids, ips, ports, initials = utils.FetchProcesses()
	totalNum := len(ids)
	var processes []application.Process
	control := make(chan bool)

	for i := 0; i < totalNum; i++ {
		process := application.Process{
			Id:    ids[i],
			Ip:    ips[i],
			Port:  ports[i],
			Ln:    nil,
			InitialV: initials[i],
			Conns: nil,
		}
		processes = append(processes, process)
	}
	processes = network.Initialize(processes)

	//test := application.Message{
	//	S: processes[0],
	//	R: 1,
	//	M: initials[0],
	//}
	//network.UnicastSend(processes[1].Conns[1], test, maxDelay, minDelay)

	start := time.Now()
	StartSimulation(processes, fNum, totalNum, maxDelay, minDelay, control)

	//for _, process := range processes {
	//	go func() {
	//		print(process.Id)
	//		messages := make(chan application.Message)
	//		firstMsg := application.Message{
	//			S: process,
	//			R: 1,
	//			M: rand.Float64(),
	//		}
	//		go network.ProcessMsg(process.Ln, messages)
	//		network.ApproximateConsensus(messages, process, fNum, totalNum, firstMsg, maxDelay, minDelay)
	//	}()
	//}

	for {
		select {
		case <-control:
			//terminate when each of them are close within 0.001
			counter++
			if counter == totalNum-fNum{
				elapsed := time.Since(start)
				log.Printf("consensus achieved using %s", elapsed)
				return
			}
		default:
			continue
		}
	}
}

func StartSimulation(processes []application.Process, fNum int, totalNum int,
	maxDelay int, minDelay int, control chan bool) {
	for _, process := range processes {
		go StartProcess(process, fNum, totalNum, maxDelay, minDelay, control)
	}
}

func StartProcess(process application.Process, fNum int, totalNum int,
	maxDelay int, minDelay int, control chan bool) {
	messages := make(chan application.Message, totalNum)
	go network.ProcessMsg(process.Ln, messages)
	network.ApproximateConsensus(messages, process, fNum, totalNum, maxDelay, minDelay, control)
}