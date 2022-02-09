package main

import (
	"fmt"
	"net"
	"tcp-server/frame"
	"tcp-server/packet"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8088")
	if err != nil {
		panic("get dial error")
	}
	id := 1
	codec := frame.MyFrameCodec{}
	connPack := packet.ConnPacket{}
	connPack.Id = string(id)
	connPack.Payload = []byte("hello world")
	b, err := connPack.Encode()
	if err != nil {
		fmt.Println("encode packet error", err)
	}
	err = codec.Encode(conn, b)
	if err != nil {
		fmt.Println("encode frame error", err)
	}
}
