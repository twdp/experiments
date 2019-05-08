package main

import (
	"context"
	"experiments/rpcx/api"
	"flag"
	"github.com/pkg/errors"
	"github.com/smallnest/rpcx/client"
	"log"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &api.Args{
		A: 10,
		B: 20,
	}

	reply := &api.Reply{}
	err := xclient.Call(context.Background(), "Mul2", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", errors.WithStack(err))
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)

}
