package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
    connHost = "localhost"
    connPort = "8088"
    connType = "tcp"
)

func main() {
    fmt.Println("Connecting to " + connType + " server " + connHost + ":" + connPort);
    conn, err := net.Dial(connType, connHost+":"+connPort)
    if err != nil {
        fmt.Println("Error connecting:", err.Error())
        os.Exit(1)
    }
    reader := bufio.NewReader(os.Stdin)
	 writer := bufio.NewWriter(conn)
    for {
        fmt.Print("Text to send: ")

        input, _ := reader.ReadString('\n')
		  writer.WriteString(input)
		  writer.Flush()

        message, _  := bufio.NewReader(conn).ReadString('\n')
        log.Print("Server relay:", message)
    }
	}