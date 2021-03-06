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

var gouroutinCount int = 10

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

	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		panic("get dial error")
	}

	codec := frame.MyFrameCodec{}
	go handleRecv(conn, codec, quit, no)

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
		fmt.Printf("[client %d]: send id=%s data=%s\n", no, submit.Id, submit.Payload)

		time.Sleep(1 * time.Second)

		if i >= 5 {
			quit <- struct{}{}
			fmt.Printf("[client %d]: exit ok\n", no)
			return
		}

	}
}

func handleRecv(conn net.Conn, codec frame.MyFrameCodec, quit chan struct{}, no int) {
	defer conn.Close()
	for {
		select {
		case <-quit:
			return

		default:

		}

		conn.SetReadDeadline(time.Now().Add(1 * time.Second))

		framePayload, err := codec.Decode(conn)
		if err != nil {
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

			fmt.Printf("[client %d]: recv connAck id=%s payload=%d\n", no, connAck.Id, connAck.Result)
		case *packet.SubmitAckPacket:

			submitAck := ack.(*packet.SubmitAckPacket)
			fmt.Printf("[client %d]: recv submtAck id=%s payload=%d\n", no, submitAck.Id, submitAck.Result)

		}

	}
}
