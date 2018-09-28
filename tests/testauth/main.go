package main

import (
	"context"
	"fmt"
	"github.com/herlegs/EtcdPlay/constant"
	"github.com/herlegs/EtcdPlay/util"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints: constant.Cluster,
		DialTimeout: time.Second * 5,
		Username: "etcdplay",
		Password:"play",
	})

	if err != nil {
		fmt.Printf("error init client: %v\n",err)
		return
	}
	defer client.Close()

	//test watch
	//channel := client.Watch(context.Background(), "test")
	//if err != nil {
	//	fmt.Printf("error get: %v\n",err)
	//}
	//
	//for {
	//	select {
	//		case resp,ok := <- channel:
	//			if !ok {
	//				fmt.Printf("channel closed\n",)
	//				return
	//			}
	//			fmt.Printf("watch event:%v\n",util.PrintWatchResponse(resp))
	//	}
	//}
	for {
		resp, err := client.Get(context.Background(), "test")
		if err != nil {
			fmt.Printf("error get: %v\n",err)
		}
		fmt.Printf("get: %v\n",util.PrintGetResponse(resp))
		time.Sleep(time.Second*10)
	}

}