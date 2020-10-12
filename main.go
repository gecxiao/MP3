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
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"./utils"
)

func main() {
	//open config file to load the processes


	round := 1

}
