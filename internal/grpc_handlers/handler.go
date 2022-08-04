package grpc_handlers

import (
	context "context"
	"service/internal/memcached_service"
	"service/pkg/cache"
)

type CacheService struct {
	memcached cache.Cache
	memcached_service.UnimplementedMemcachedServiceServer
}

func NewCacheService(memcached cache.Cache) *CacheService {
	return &CacheService{memcached: memcached}
}

func (c CacheService) Get(ctx context.Context, in *memcached_service.GetRequest) (*memcached_service.GetResponse, error) {
	val, err := c.memcached.Get(in.GetKey())
	if err != nil {
		return &memcached_service.GetResponse{
			Success: false,
			Error:   err.Error(),
			Item:    nil,
		}, err
	}

	return &memcached_service.GetResponse{
		Success: true,
		Error:   "",
		Item: &memcached_service.Item{
			Key:   in.Key,
			Value: val,
		},
	}, nil
}

func (c CacheService) Set(ctx context.Context, in *memcached_service.SetRequest) (*memcached_service.SetResponse, error) {
	err := c.memcached.Set(in.GetKey(), in.GetValue())
	if err != nil {
		return &memcached_service.SetResponse{
			Success: false,
			Error:   err.Error(),
		}, err
	}

	return &memcached_service.SetResponse{
		Success: true,
		Error:   "",
	}, nil
}

func (c CacheService) Delete(ctx context.Context, in *memcached_service.DeleteRequest) (*memcached_service.DeleteResponse, error) {
	err := c.memcached.Delete(in.GetKey())
	if err != nil {
		return &memcached_service.DeleteResponse{
			Success: false,
			Error:   err.Error(),
		}, err
	}

	return &memcached_service.DeleteResponse{
		Success: true,
		Error:   "",
	}, nil
}
