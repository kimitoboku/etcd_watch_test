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
	etcdWatchKey := flag.String("etcdWatchKey", "foo", "etcd key to watch")
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

	fmt.Println("started goroutine for PUT...")
	for {
		fmt.Println("hogehoge")
		_, err := etcd.Put(context.TODO(), *etcdWatchKey, time.Now().String())
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("populated " + *etcdWatchKey + " with a value..")
		time.Sleep(2 * time.Second)
	}

}
