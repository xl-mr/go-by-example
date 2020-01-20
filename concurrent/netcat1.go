package main

import "net"

import "flag"

import "log"

import "os"

import "io"

func main() {
	port := flag.String("port", "8000", "")
	conn, err := net.Dial("tcp", ":"+*port)

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
