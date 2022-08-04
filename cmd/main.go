package main

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"service/internal/grpc_handlers"
	"service/internal/memcached_service"
	"service/pkg/cache"
	"time"
)

func main() {
	serv := grpc.NewServer(
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				Time:    30 * time.Second,
				Timeout: 5 * time.Second,
			},
		),
	)
	memcached, err := cache.NewMemcached(memcache.New("127.0.0.1:11211"))
	if err != nil {
		panic(err)
	}

	cs := grpc_handlers.NewCacheService(memcached)

	memcached_service.RegisterMemcachedServiceServer(serv, cs)

	fmt.Println("Listen :8080")

	listener, err := net.Listen("tcp", ":8080")

	err = serv.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}
}
