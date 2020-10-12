package application

import (
	"net"
)

type Process struct {
	Id    string
	Ip    string
	Port  string
	Ln    net.Listener
	Conns []net.Conn
	InitialV float64
}

type Message struct {
	S string
	R int
	M float64
}

