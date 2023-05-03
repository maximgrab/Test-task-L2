package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	var timeout string
	flag.StringVar(&timeout, "timeout", "10s", "connection establishing timeout")
	flag.Parse()
	host := flag.Arg(0)
	port := flag.Arg(1)

	timeoutInt, err := strconv.Atoi((timeout[:len(timeout)-1]))
	if err != nil {
		log.Println(err)
	}
	var connection net.Conn
	start := time.Now()
	for time.Since(start) < (time.Duration(timeoutInt) * time.Second) {
		connection, _ = net.Dial("tcp", host+":"+port)
	}
	defer connection.Close()
	go func() {
		reader := bufio.NewReader(connection)
		for {
			msg, err := reader.ReadString('\n')
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Println(msg)
		}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := scanner.Text()
		_, err := fmt.Fprintf(connection, data+"\n")
		if err != nil {
			log.Println(err)
		}
	}
}
