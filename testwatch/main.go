package main

import (
	"context"
	"fmt"
	"github.com/herlegs/EtcdPlay/util"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	etcdHost := "127.0.0.1:2379"
	etcdWatchKey := "test"

	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"http://" + etcdHost, /*"http://127.0.0.1:22379"*/},
		DialTimeout: time.Second * 5,
	})

	if err != nil {
		fmt.Printf("error init client: %v\n",err)
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
			case resp, ok := <- watchCh:
				if ok {
					for _, event := range resp.Events {
						fmt.Printf("Event received[%s]! %s executed on %q with value %q\n",util.Parse(resp.Header.MemberId), event.Type, event.Kv.Key, event.Kv.Value)
					}
				} else {
					fmt.Printf("channel closed\n",)
					return
				}
		}
	}
}

