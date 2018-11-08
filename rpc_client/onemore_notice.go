package rpc_client

import (
	"time"

	"github.com/hprose/hprose-golang/rpc"
)

type OnemoreServer struct {
	PostSystem func(value string) error
}

func OnemoreNotice(value string, url string) {
	var onemoreServer *OnemoreServer
	client := rpc.NewHTTPClient(url)
	client.UseService(&onemoreServer)
	client.SetTimeout(120 * time.Second)

	onemoreServer.PostSystem(value)
}
