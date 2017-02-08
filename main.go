package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	fmt.Println("Socket sender")

	fmt.Println("Start TCP client")

	// open a socket to the listener
	conn, err := net.Dial("tcp", "127.0.0.1:4000")
	if err != nil {
		fmt.Println("Unable to open TCP port")
		return
	}

	buf := make([]byte, 1024)

	now := time.Now().UTC()

	var count = 0
	for {

		_, err := conn.Write(buf)
		if err != nil {
			fmt.Println("Write error")
			break
		}

		count++

		if time.Since(now) >= 1000000000 {
			break
		}
	}

	fmt.Println("count = ", count)
	conn.Close()
}
