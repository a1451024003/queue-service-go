package main

import (
	"fmt"
	"io/ioutil"
	"queue-service-go/rpc_client"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/yaml.v2"
)

var config Config

type Config struct {
	Rpc                  string `yaml:"rpc,omitempty"`
	RpcGroup             string `yaml:"rpc_group,omitempty"`
	RpcCourseware        string `yaml:"rpc_courseware,omitempty"`
	RpcGroupSubscribes   string `yaml:"rpc_group_subscribes,omitempty"`
	RpcOnemorePushNotice string `yaml:"rpc_onemore_notice,omitempty"`
	Redis                Redis  `yaml:"redis,omitempty"`
}

type Redis struct {
	Host                string `yaml:"host,omitempty"`
	Port                string `yaml:"port,omitempty"`
	SxVoteListName      string `yaml:"sx_vote_list_name,omitempty"`
	GroupListName       string `yaml:"group_list_name,omitempty"`
	PackTaskName        string `yaml:"packtask_name,omitempty"`
	OnemorePushNotice   string `yaml:"onemore_notice_name,omitempty"`
	GroupSubscribesName string `yaml:"group_subscribes_name,omitempty"`
}

func createConfig() {
	var conf Config
	yamlFile, err := ioutil.ReadFile("./conf/conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(conf)

	config = conf
}

func init() {
	createConfig()
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			logs.SetLogger(logs.AdapterFile, `{"filename":"queue.log"}`)
			logs.Error(err)
		}
	}()
	redisConf := fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port)
	redisConn, _ := redis.Dial("tcp", redisConf)
	for {
		if reply, err := redis.Values(redisConn.Do("BRPOP", config.Redis.GroupListName, config.Redis.PackTaskName, config.Redis.SxVoteListName, config.Redis.GroupSubscribesName, config.Redis.OnemorePushNotice, 0)); err == nil {
			key := string(reply[0].([]byte))
			value := string(reply[1].([]byte))
			fmt.Println(key, value)

			if key == config.Redis.GroupListName {
				rpc_client.GroupCoursewarePush(value, config.RpcGroup) //圈子服务：课件推送队列
			} else if key == config.Redis.SxVoteListName {
				rpc_client.ActivitySXVote(value, config.Rpc) //活动服务：山西活动_投票队列
			} else if key == config.Redis.PackTaskName {
				rpc_client.CoursewarePackTask(value, config.RpcCourseware)
			} else if key == config.Redis.GroupSubscribesName {
				rpc_client.GroupSubscribe(value, config.RpcGroupSubscribes) //圈子服务：订阅队列
			} else if key == config.Redis.OnemorePushNotice {
				rpc_client.OnemoreNotice(value, config.RpcOnemorePushNotice) //万摩服务：发布通知队列
			} else {
				continue
			}
		} else {
			time.Sleep(5 * time.Second)
			continue
		}
	}
}
