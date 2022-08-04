package cache

type InMemory struct {
	storage map[string][]byte
}

func NewInMemory() *InMemory {
	storage := make(map[string][]byte)

	return &InMemory{storage: storage}
}

//Get Return value from the store
func (i *InMemory) Get(key string) ([]byte, error) {
	val, ok := i.storage[key]
	if !ok {
		return nil, ErrKeyNotExist
	}
	return val, nil
}

//Set Writes a value by key to the store
func (i *InMemory) Set(key string, value []byte) error {
	i.storage[key] = value
	return nil
}

//Delete Deletes a value by key to the store
func (i *InMemory) Delete(key string) error {
	_, err := i.Get(key)
	if err != nil {
		return err
	}
	delete(i.storage, key)
	return nil
}
