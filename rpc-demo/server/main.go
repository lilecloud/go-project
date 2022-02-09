package main

import (
	"net"
	"net/rpc"
)

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, _ := net.Listen("tcp", ":123")

	for {
		conn, _ := listener.Accept()

		go rpc.ServeConn(conn)
	}

}

type HelloService struct{}

func (hello *HelloService) Say(req string, reply *string) error {
	*reply = ("hello " + req)
	return nil
}
