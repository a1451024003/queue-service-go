package rpc_client

import (
	"github.com/hprose/hprose-golang/rpc"
	"time"
)

type GroupServer struct {
	Subscribe func(string)
}

func GroupSubscribe(value string, url string) {
	var groupServer *GroupServer
	client := rpc.NewHTTPClient(url)
	client.UseService(&groupServer)
	client.SetTimeout(120 * time.Second)

	groupServer.Subscribe(value)
}
