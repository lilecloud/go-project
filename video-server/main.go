package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {

	listener, _ := net.Listen("tcp", ":9090")

	for {
		conn, _ := listener.Accept()
		fmt.Println("receiveConn")
		bytes, _ := ioutil.ReadAll(conn)
		fmt.Println(string(bytes))
		conn.Write([]byte("hello world"))
	}
}
