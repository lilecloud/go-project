package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		panic("listen 127.0.0.1:9999 error ")
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}

		go process(conn)
	}

}

func process(conn net.Conn) {

	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)

		var bytes [128]byte

		n, err := reader.Read(bytes[:])
		if err != nil {
			fmt.Println("read error")
			return
		}
		receStr := string(bytes[:n])
		fmt.Println("receive message from client:********", receStr, "*********************")
		conn.Write([]byte(receStr))
	}

}
