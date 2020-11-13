package main

import (
	"context"
	"fmt"
	"github.com/smallnest/rpcx/client"
	"log"
)

type args struct {
	A int
	B int
}

type reply struct {
	C int
}

func main() {
	d := client.NewPeer2PeerDiscovery("tcp@localhost:8972", "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	a := &args{
		3,
		2,
	}

	r := &reply{}

	err := xclient.Call(context.Background(), "Mul", a, r)
	if err != nil {
		fmt.Println("errors in Call", err)
	}

	log.Printf("%d * %d = %d", a.A, a.B, r.C)
}
