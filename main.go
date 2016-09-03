// Package main is a simple server that echos back to a client whatever is sent
// and then closes the connection.
package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {
	listenAddr := ":8080"
	if envAddr := os.Getenv("LISTEN_PORT"); envAddr != "" {
		listenAddr = envAddr
	}

	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("could not open TCP listener: %s", err.Error())
		return
	}

	log.Printf("waiting for connections on port %s\n", listenAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("could not accept connection: %s", err.Error())
		}
		log.Printf("connection accepted from %s <-> %s\n", conn.RemoteAddr(), conn.LocalAddr())

		b := bufio.NewReader(conn)
		line, err := b.ReadBytes('\n')
		if err != nil {
			log.Fatalf("error writing to connection: %s", err.Error())
		}

		conn.Write(line)
		conn.Close()
	}
}
