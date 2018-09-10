package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	etcdHost := "etcd-int.stg-myteksi.com:2379"
	etcdWatchKey := "csdp:v2:crudds"

	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"http://" + etcdHost},
		DialTimeout: time.Second * 5,
		Username:"gargamel",
		Password: "ciuBLuDrX6",
	})

	if err != nil {
		fmt.Printf("error init client: %v\n",err)
		return
	}
	defer client.Close()

	resp, err := client.Get(context.Background(), etcdWatchKey, clientv3.WithPrefix())
	if err != nil {
		fmt.Printf("err when get: %v\n",err)
	}
	oneKey := resp.Kvs[1]
	fmt.Printf("rev[%v] key[%v] version[%v] createRevision[%v] modRevision[%v]\n",resp.Header.Revision,
		string(oneKey.Key),oneKey.Version, oneKey.CreateRevision, oneKey.ModRevision)

	watchCh := client.Watch(context.Background(), etcdWatchKey,
		clientv3.WithPrefix(),
		clientv3.WithRev(resp.Header.Revision))

	for resp := range watchCh {
		for _, event := range resp.Events {
			fmt.Printf("Event received! %s executed on %q with value %q\n", event.Type, event.Kv.Key, event.Kv.Value)
		}
	}
}
