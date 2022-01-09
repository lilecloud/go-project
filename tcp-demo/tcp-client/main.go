package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		panic("get dial error")
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Scanln(a ...interface{})
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			panic("read line error")

		}

		_, err1 := conn.Write([]byte(str))
		if err1 != nil {
			panic("write err" + err1.Error())
		}

		buf := [512]byte{}

		n, err := conn.Read(buf[:])

		fmt.Println("receive messag from server:", string(buf[:n]))

	}
}
