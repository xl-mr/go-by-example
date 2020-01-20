package main

import "net"

import "log"

import "io"

import "time"

import "flag"

import "fmt"

var port *string = flag.String("port", "8000", "")

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	for {
		str := fmt.Sprintf("port : %s, time : %s\n", *port, time.Now().Format("15:04:05"))
		_, err := io.WriteString(c, str)
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
