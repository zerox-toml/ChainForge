package main

import (
	"bufio"
	"fmt"
	"os"
)

func readPorts(myPort string) []string {
	var ports []string

	readFile, err := os.Open("ports/ports.txt")
	if err != nil {
		os.Exit(0)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		ports = append(ports, fileScanner.Text())
	}
	readFile.Close()
	return ports
}

func checkIfInList(port string, ports []string) bool {
	for i, cPort := range ports {
		fmt.Println(i, cPort, port)
		if cPort == port {
			return false
		}
	}
	return true
}

func appendPort(port string, ports []string) {
	if !checkIfInList(port, ports) {
		fmt.Println("Here")
		return
	}
	f, err := os.OpenFile("ports/ports.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(port + "\n"); err != nil {
		panic(err)
	}
}
