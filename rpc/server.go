package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Listener int

type Reply struct {
	Data string
}

func (l *Listener) GetLine(line []byte, reply *Reply) error {
	rv := string(line)
	fmt.Printf("Receive: %v\n", rv)
	*reply = Reply{rv}
	return nil
}

func main() {
	inbound, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal(err)
	}

	listener := new(Listener)
	_ = rpc.Register(listener)
	rpc.Accept(inbound)
}
