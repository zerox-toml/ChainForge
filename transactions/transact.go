package main

import (
	"fmt"
	"net"
)

const (
	SERVER_TYPE = "tcp"
)

func main() {
	// dial("9988")
	// str := "[BROADCAST]9988"

	// blunt := str[11:]
	// fmt.Println(blunt)
	// if blunt == "[BROADCAST]" {
	// }
	dial("9988")
}

func dial(port string) (net.Conn, error) {
	fmt.Println("dialing")
	connection, err := net.Dial(SERVER_TYPE, "localhost:"+port)
	if err != nil {
		return nil, err
	}
	_, err = connection.Write([]byte("[TRANSACTION]0x0000,0x0066,50"))
	buffer := make([]byte, 1024)
	messageLength, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:messageLength]))
	return connection, nil
}
