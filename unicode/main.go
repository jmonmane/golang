package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fd, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		fmt.Printf("%x\n", scanner.Bytes())
	}

}
