package main

import (
	"fmt"
	"net"
	"sync"
	"tcp-server/frame"
	"tcp-server/packet"
	"time"

	"github.com/lucasepe/codename"
)

var gouroutinCount int = 2

func main() {
	group := sync.WaitGroup{}
	group.Add(gouroutinCount)

	for i := 1; i <= gouroutinCount; i++ {
		go func(i int) {
			startClient(&group, i)
		}(i)
	}
	group.Wait()
}

func startClient(group *sync.WaitGroup, no int) {
	defer group.Done()

	quit := make(chan struct{})

	conn, err := net.Dial("tcp", "127.0.0.1:9089")
	if err != nil {
		panic("get dial error")
	}

	codec := frame.MyFrameCodec{}
	go handleRecv(conn, codec, quit)

	rng, _ := codename.DefaultRNG()

	for i := 1; ; i++ {
		submit := packet.SubmitPacket{
			Id: fmt.Sprintf("%08d", i),
		}
		data := codename.Generate(rng, i)
		submit.Payload = []byte(data)

		bytes, err := packet.Encode(&submit)
		if err != nil {
			fmt.Println(err)
			continue
		}
		er1r := codec.Encode(conn, bytes)
		if er1r != nil {
			fmt.Println(er1r)
			continue
		}
		fmt.Printf("send data=%s\n", data)

		if i >= 5 {
			quit <- struct{}{}
			fmt.Printf("[client %d]: exit ok\n", i)
			return
		}

	}
}

func handleRecv(conn net.Conn, codec frame.MyFrameCodec, quit chan struct{}) {
	defer conn.Close()
	for {
		select {
		case <-quit:
			return

		default:

		}

		conn.SetReadDeadline(time.Now().Add(2 * time.Second))

		framePayload, err := codec.Decode(conn)
		if err != nil {
			fmt.Println("client codec decode err", err)
			if e, ok := err.(net.Error); ok {
				if e.Timeout() {
					continue
				}
			}
			panic(err)

		}
		ack, err := packet.Decode(framePayload)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch ack.(type) {
		case *packet.ConnAckPacket:
			connAck := ack.(*packet.ConnAckPacket)

			fmt.Printf("recv connAck id=%s payload=%d\n", connAck.Id, connAck.Result)
		case *packet.SubmitAckPacket:

			submitAck := ack.(*packet.SubmitAckPacket)
			fmt.Printf("recv submtAck id=%s payload=%d\n", submitAck.Id, submitAck.Result)

		}

	}
}
