package utils

import (
	"os"
	"bufio"
	"strconv"
	"strings"
	"log"
	//"fmt"
)

//Fetches all the ports
func FetchPorts()([]string, []string, []string){
	line := 0
	f, err := os.Open("./config.txt")
	var ids []string
	var ips []string
	var ports []string
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if(line != 0) {
			id := strings.Split(scanner.Text(), " ")[0]
			ip := strings.Split(scanner.Text(), " ")[1]
			port := strings.Split(scanner.Text(), " ")[2]
			ids = append(ids, id)
			ips = append(ips, ip)
			ports = append(ports, port)
		}
		line = line + 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return ids, ips, ports
}



//parses the min and max delays from the config file
func FetchDelay()(int, int, int){
	f, err := os.Open("./config.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	delays := strings.Fields(scanner.Text())
	min_delay, _ := strconv.Atoi(delays[0])
	max_delay, _ := strconv.Atoi(delays[1])
	fNum, _ :=strconv.Atoi(delays[2])
	f.Close()
	return min_delay, max_delay, fNum
}