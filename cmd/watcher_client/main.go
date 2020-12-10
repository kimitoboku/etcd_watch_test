package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	etcdHost := flag.String("etcdHost", "0.0.0.0:2379", "etcd host")
	etcdWatchKey := flag.String("etcdWatchKey", "watch_prefix", "etcd key to watch")
	flag.Parse()

	fmt.Println("connecting to etcd - " + *etcdHost)

	etcd, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://" + *etcdHost},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("connected to etcd - " + *etcdHost)

	defer etcd.Close()

	watchChan := etcd.Watch(context.Background(), *etcdWatchKey)
	fmt.Println("set WATCH on " + *etcdWatchKey)

	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			fmt.Printf("Received Now %s, %s executed on %q with value %q\n", time.Now().String() ,event.Type, event.Kv.Key, event.Kv.Value)
		}
	}
}
