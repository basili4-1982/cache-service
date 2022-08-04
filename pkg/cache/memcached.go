package cache

import (
	"errors"
	"github.com/bradfitz/gomemcache/memcache"
)

var ErrReqCli = errors.New("memcached client needs to be installed")

type MemcachedClient interface {
	Set(item *memcache.Item) error
	Get(key string) (item *memcache.Item, err error)
	Delete(key string) error
}

type Memcached struct {
	client MemcachedClient
}

func NewMemcached(client MemcachedClient) (*Memcached, error) {
	if client == nil {
		return nil, ErrReqCli
	}
	return &Memcached{client: client}, nil
}

func (m Memcached) Get(key string) ([]byte, error) {
	val, err := m.client.Get(key)
	if err != nil {
		if err == memcache.ErrMalformedKey {
			return nil, ErrKeyNotExist
		}

		return nil, err
	}

	return val.Value, nil
}

func (m Memcached) Set(key string, value []byte) error {
	err := m.client.Set(&memcache.Item{
		Key:        key,
		Value:      value,
		Flags:      0,
		Expiration: 0,
	})
	if err != nil {
		return err
	}
	return nil
}

func (m Memcached) Delete(key string) error {
	return m.client.Delete(key)
}
