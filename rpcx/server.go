package main

import (
	"context"
	"github.com/smallnest/rpcx/server"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

type Arith int

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

func main() {
	s := server.NewServer()
	_ = s.RegisterName("Arith", new(Arith), "")
	_ = s.Serve("tcp", ":8972")
}