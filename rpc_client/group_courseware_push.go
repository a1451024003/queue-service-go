package rpc_client

import (
	"github.com/hprose/hprose-golang/rpc"
	"time"
)

type NewsServer struct {
	PushCourseware func(string)
}

func GroupCoursewarePush(value string, url string) {
	var newsServer *NewsServer
	client := rpc.NewHTTPClient(url)
	client.UseService(&newsServer)
	client.SetTimeout(120 * time.Second)

	newsServer.PushCourseware(value)
}
