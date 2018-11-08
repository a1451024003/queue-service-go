package rpc_client

import (
	"github.com/hprose/hprose-golang/rpc"
	"strconv"
)

type CoursewareServer struct {
	PackTask func(int) interface{}
}

func CoursewarePackTask(value string, url string) {
	catgory_id, _ := strconv.Atoi(value)
	var packTaskServer *CoursewareServer
	client := rpc.NewHTTPClient(url)
	client.UseService(&packTaskServer)
	packTaskServer.PackTask(catgory_id)
}
