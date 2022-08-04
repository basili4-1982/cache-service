package cache

import "errors"

var ErrKeyNotExist = errors.New("key not found")

type Cache interface {
	//Get Return value from the store
	Get(key string) ([]byte, error)
	//Set Writes a value by key to the store
	Set(key string, value []byte) error
	//Delete Deletes a value by key to the store
	Delete(key string) error
}
