package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, _ := rpc.Dial("tcp", "localhost:123")

	var reply string = ""
	err := client.Call("HelloService.Say", "golang", &reply)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply)

}
