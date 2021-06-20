package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	grantResp, err := cli.Grant(context.TODO(), 10)
	if err != nil {
		log.Fatal(err)
	}

	key := "foo"
	_, err = cli.Put(context.TODO(), key, "bar", clientv3.WithLease(grantResp.ID))
	if err != nil {
		log.Fatal(err)
	}

	// the key "foo" will be kept forever
	chanKeepAliveResp, err := cli.KeepAlive(context.TODO(), grantResp.ID)
	if err != nil {
		log.Fatal(err)
	}

	// watch keepAlive status
	done := make(chan struct{})
	go func() {
		// The returned "LeaseKeepAliveResponse" channel closes if underlying keep
		// alive stream is interrupted in some way the client cannot handle itself;
		// given context "ctx" is canceled or timed out.
		for keepAliveResp := range chanKeepAliveResp {
			log.Printf("ttl for key %s: %v", key, keepAliveResp.TTL)

		}
		select {
		case <-done:
		default:
			err := errors.New(fmt.Sprintf("the underlying keep alive stream for key %s is interrupted in some way the client cannot handle itself, e.g. context \"ctx\" is canceled or timed out", key))
			panic(err)
		}
	}()

	// sleep 10 seconds to simulate program running
	time.Sleep(20 * time.Second)
	// now it's time to end this program
	close(done)
}

// How does keepAlive determine the interval between two renewals ?
// $ go run Lease_With_KeepAlive.go
// 2021/06/20 10:00:40 ttl for key foo: 10
// 2021/06/20 10:00:43 ttl for key foo: 10
// 2021/06/20 10:00:47 ttl for key foo: 10
// 2021/06/20 10:00:50 ttl for key foo: 10
// 2021/06/20 10:00:54 ttl for key foo: 10
// 2021/06/20 10:00:57 ttl for key foo: 10
