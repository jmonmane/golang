package main

import (
	"bufio"
	"fmt"
	"net"
)

const (
	target = "45.79.93.134"
)

func connect() {
	c, err := net.Dial("tcp", target+":22")
	if err != nil {
		panic(err)
	}
	defer c.Close()
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func listen() {
	l, err := net.Listen("tcp", ":9001")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		fmt.Printf("%s", "We received an incoming connection")
		s := []byte{}
		conn.Read(s)

	}
}
func main() {
	go connect()
	go listen()
	c, _ := net.Dial("tcp", "127.0.0.1:9001")
	defer c.Close()
	s := "text to send"
	c.Write([]byte(s))
}
