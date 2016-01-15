package main

import (
	"io"
	"log"
	"net"
	"os"
	"bufio"
	"strings"
)

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	i := 0
	for i < 10 {
		scanner.Scan()
		ln := scanner.Text()
		if strings.Split(ln, ":")[0] == "User-Agent" {
			io.WriteString(conn, strings.Split(ln, ":")[1])
			break	
		} else {
				i++
			}
	}
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
