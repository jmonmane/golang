package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func handle(conn net.Conn) {
	ra := conn.RemoteAddr().String()
	both := strings.Split(ra, ":")
	io.WriteString(conn, both[0])
	conn.Close()
	return
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Must provide port number")
	}
	port := os.Args[1]
	server, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err)
	}
	defer server.Close()

	for {
		conn, _ := server.Accept()
		go handle(conn)
	}

}
