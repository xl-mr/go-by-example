package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main()  {
	for _, val := range os.Args[1:] {
		addr := strings.SplitN(val, "=", 2)
		go conTcp(addr[1])
	}
	for {
		time.Sleep(1 * time.Second)
	}
}

func conTcp(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for {
		mustCopy1(os.Stdout, conn)
		time.Sleep(1 * time.Second)
	}
}

func mustCopy1(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}