package main

import (
	"context"
	"fmt"
	"time"

	"github.com/herlegs/EtcdPlay/constant"
	"github.com/herlegs/EtcdPlay/util"
	"go.etcd.io/etcd/clientv3"
)

func main() {
	etcdHost := constant.Host1
	etcdWatchKey := "test"

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{etcdHost},
		DialTimeout: time.Second * 5,
	})

	if err != nil {
		fmt.Printf("error init client: %v\n", err)
		return
	}
	defer client.Close()

	//resp, err := client.Get(context.Background(), etcdWatchKey, clientv3.WithPrefix())
	//if err != nil {
	//	fmt.Printf("err when get: %v\n",err)
	//}
	//oneKey := resp.Kvs[1]
	//fmt.Printf("revison[%v] key[%v] version[%v] createRevision[%v] modRevision[%v]\n",resp.Header.Revision,
	//	string(oneKey.Key),oneKey.Version, oneKey.CreateRevision, oneKey.ModRevision)

	watchCh := client.Watch(context.Background(), etcdWatchKey,
		clientv3.WithPrefix(),
		clientv3.WithRev(0))

	for {
		select {
		case resp, ok := <-watchCh:
			if ok {
				for _, event := range resp.Events {
					fmt.Printf("Event received[%s]! %s executed on %q with value %q\n", util.GetHostFromMemberUint64ID(resp.Header.MemberId), event.Type, event.Kv.Key, event.Kv.Value)
				}
			} else {
				fmt.Printf("channel closed\n")
				return
			}
		}
	}
}
