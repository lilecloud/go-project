package main

import (
	"fmt"
	"net"
	"tcp-server/frame"
	"tcp-server/packet"
)

func main() {
	listener, err := net.Listen("tcp", ":8088")
	if err != nil {
		fmt.Println("lisent port err", err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept connect err", err)
			return
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	frameCodec := frame.MyFrameCodec{}
	for {
		payload, err := frameCodec.Decode(conn)
		if err != nil {
			fmt.Println("decode frame payload error", err)
			return
		}
		handlePacket(payload)
	}

}

func handlePacket(b []byte) ([]byte, error) {
	p, err := packet.Decode(b)
	if err != nil {
		fmt.Println("decode packet error", err)
		return nil, err
	}
	switch p.(type) {
	case *packet.ConnPacket:
		connPack := p.(*packet.ConnPacket)
		id := connPack.Id

		ackPack := packet.ConnAckPacket{}
		ackPack.Id = id
		ackPack.Result = 1
		return ackPack.Encode()
	}

	return nil, nil
}
