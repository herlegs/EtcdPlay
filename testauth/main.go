package main

import (
	"context"
	"fmt"
	"github.com/herlegs/EtcdPlay/constant"
	"github.com/herlegs/EtcdPlay/util"
	"time"
	"go.etcd.io/etcd/clientv3"
)

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{constant.Host1},
		DialTimeout: time.Second * 5,
		Username: "etcdplay",
		Password:"play",
	})

	if err != nil {
		fmt.Printf("error init client: %v\n",err)
		return
	}
	defer client.Close()

	resp, err := client.Get(context.Background(), "test")
	if err != nil {
		fmt.Printf("error get: %v\n",err)
	} else {
		fmt.Printf("res:%v\n",util.PrintGetResponse(resp))
	}


}