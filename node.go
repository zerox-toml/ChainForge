package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type wallet struct {
	address string
	funds   float32
}

// This is a comment:
const (
	SERVER_HOST = "localhost" // comment
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main() {
	port := getPort()
	myPort := strconv.Itoa(port)
	ports := readPorts(myPort)
	server, err := net.Listen(SERVER_TYPE, "localhost:"+myPort)

	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer server.Close()
	broadcastEveryNode(ports, myPort)
	for {
		listenConnections(server, myPort, ports)
	}
}

// here we should get a transaction
func listenConnections(server net.Listener, myPort string, ports []string) {
	fmt.Println("Listening on " + "localhost:" + myPort)
	fmt.Println("Waiting for client...")
	connection, err := server.Accept()
	if err != nil {
		fmt.Println("Error accepting: ", err.Error())
		os.Exit(1)
	}
	fmt.Println("client connected")
	go processClient(connection, myPort, ports)
}

// this one is called when node comes online
func broadcastEveryNode(ports []string, myPort string) {
	for _, c_port := range ports {
		if c_port != myPort { // only connect to the ports that are not the nodes port
			dialConn, err := dial("[BROADCAST]", c_port)
			if err != nil {
				fmt.Println("Not able to connect on :", c_port)
			} else {
				dialConn.Close()
			}
		}
	}
}

// here we handle the transaction
func processClient(connection net.Conn, port string, ports []string) {

	buffer := make([]byte, 1024)
	messageLength, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// we get the value
	received := string(buffer[:messageLength])
	// fmt.Println("Received: ", received[:13])
	// fmt.Println("HERHEHRHERHERHEHHERHEH")
	if received[:11] == "[BROADCAST]" {
		appendPort(received[11:], ports)
		// add functionality here that broadcasts the new online node to other nodes?
	} else if received[:13] == "[TRANSACTION]" {
		processTransaction(connection, received)
	}
	connection.Close()
}

func transactionResponse(connection net.Conn, message string) {
	_, err := connection.Write([]byte(string(message)))
	if err != nil {
		fmt.Print("TRANSACTION PROBLEM IN WRITE: ")
		panic(err)
	}
}

func processTransaction(connection net.Conn, received string) {
	transactionResult, toAddr, fromAddr, fromBalance, toBalance := transactionValidation(received[13:])

	fmt.Println(transactionResult)
	switch transactionResult {
	case "OK":
		updateBalance(fromAddr, strconv.Itoa(fromBalance))
		updateBalance(toAddr, strconv.Itoa(toBalance))
	case "ERROR":
		transactionResponse(connection, "[TRANSACTION ERROR]")
	}
}

func dial(dialType string, port string) (net.Conn, error) {
	fmt.Println("dialing")
	connection, err := net.Dial(SERVER_TYPE, "localhost:"+port)
	if err != nil {
		return nil, err
	}
	_, err = connection.Write([]byte(dialType + port))
	buffer := make([]byte, 1024)
	messageLength, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:messageLength]))
	return connection, nil
}

func getPort() int {
	fmt.Println("ENTER PORT NB: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(0)
	}
	input = strings.TrimSpace(input)
	i, _ := strconv.Atoi(input)
	return i
}
