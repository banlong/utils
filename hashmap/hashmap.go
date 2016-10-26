package hashmap

import (
	"sync"
	"errors"
)

var (
	standardMap = NewHashMap()
)

type HashMap struct{
	hashmap     map[interface{}]interface{}
	lock        *sync.RWMutex
}

func NewHashMap() *HashMap {
	var ret = HashMap{
		hashmap: make(map[interface{}]interface{}),
		lock: new(sync.RWMutex),
	}
	return &ret
}

func (m *HashMap)Put(key interface{}, value interface{})  error{
	m.lock.Lock()
	defer m.lock.Unlock()
	m.hashmap[key] = value
	return nil
}


func (m *HashMap) Get(key interface{}) (interface{}, error)  {
	m.lock.RLock()
	defer m.lock.RUnlock()
	value, isOk := m.hashmap[key]
	if !isOk {
		return nil, errors.New("Item not found")
	}

	return value, nil

}

func (m *HashMap) GetSize()  int  {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return len(m.hashmap)
}


func Set(key interface{}, value interface{})  error{
	standardMap.lock.Lock()
	defer standardMap.lock.Unlock()
	standardMap.hashmap[key] = value
	return nil
}


func Get(key interface{}) (interface{}, error)  {
	standardMap.lock.RLock()
	defer standardMap.lock.RUnlock()
	value, isOk := standardMap.hashmap[key]
	if !isOk {
		return nil, errors.New("Item not found")
	}
	return value, nil

}

func Delete(key interface{})  error  {
	standardMap.lock.Lock()
	defer standardMap.lock.Unlock()
	_, isOk := standardMap.hashmap[key]
	if !isOk {
		return errors.New("Item not found")
	}else{
		delete(standardMap.hashmap, key)
		return nil
	}
}

func GetSize()  int  {
	standardMap.lock.RLock()
	defer standardMap.lock.RUnlock()
	return len(standardMap.hashmap)
}