package main

import (
	"context"
	"errors"
	"experiments/rpcx/api"
	"flag"
	"fmt"
	"github.com/smallnest/rpcx/server"
)

var (
	addr1 = flag.String("addr", "localhost:8972", "server address")
)

type Arith struct{}

// the second parameter is not a pointer
func (t *Arith) Mul(ctx context.Context, args *api.Args, reply *api.Reply) error {
	reply.C = args.A * args.B
	return nil
}

func (t *Arith) Mul2(ctx context.Context, args *api.Args, reply *api.Reply) error {
	reply.C = args.A * args.B
	fmt.Println(args.AAA.X)
	fmt.Println("===")
	return errors.New("22")
}

func main() {
	flag.Parse()

	s := server.NewServer()
	//s.Register(new(Arith), "")
	s.RegisterName("Arith", new(Arith), "")
	err := s.Serve("tcp", *addr1)
	if err != nil {
		panic(err)
	}
}
