package main

import (
	"context"
	"fmt"
	"github.com/herlegs/EtcdPlay/util"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main()  {
	etcdHost := "127.0.0.1:2379"
	//etcdWatchKey := "test"

	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"http://" + etcdHost, "http://127.0.0.1:22379","http://127.0.0.1:32379"},
		DialTimeout: time.Second * 5,
	})

	if err != nil {
		fmt.Printf("error init client: %v\n",err)
		return
	}
	defer client.Close()

	ctx,_ := context.WithTimeout(context.Background(), time.Second*3)
	resp,err := client.Grant(ctx, 30)
	if err != nil {
		fmt.Printf("error granting lease: %v\n",err)
		return
	}

	fmt.Printf("lease granted: %x\n",resp.ID)

	ch,err := client.KeepAlive(context.Background(), resp.ID)
	if err != nil {
		fmt.Printf("error keep alive\n",)
	}

	for {
		select {
			case resp,ok := <-ch:
				if ok {
					fmt.Printf("member[%s] lease[%x] ttl[%v] timeNow[%v]\n",util.GetHostFromMemberUint64ID(resp.MemberId), resp.ID, resp.TTL, time.Now().String())
				} else {
					fmt.Printf("channel closed[%v]\n",time.Now().String())
					return
				}
		}
	}

}
