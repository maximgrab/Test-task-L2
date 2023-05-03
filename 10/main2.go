package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8083")
	if err != nil {
		log.Println(err)
	}
	connection, _ := ln.Accept()
	for {
		msg, err := bufio.NewReader(connection).ReadString('\n')
		if err == io.EOF {
			return
		}
		fmt.Println(msg)
		msg = "socket " + msg
		connection.Write([]byte(msg + "\n"))
	}
}
