package rpc_client

import (
	"github.com/hprose/hprose-golang/rpc"
	"time"
)

type VoteServer struct {
	Vote func(string)
}

func ActivitySXVote(value string, url string) {
	var voteServer *VoteServer
	client := rpc.NewHTTPClient(url)
	client.UseService(&voteServer)
	client.SetTimeout(120 * time.Second)

	voteServer.Vote(value)
}
