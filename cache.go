package form

import (
	"reflect"
	"sync"
)

type cachedField struct {
	idx  int
	name string
}

type cachedStruct struct {
	fields []cachedField
}

type structCacheMap struct {
	lock sync.RWMutex
	m    map[reflect.Type]*cachedStruct
}

func (s *structCacheMap) Get(key reflect.Type) (*cachedStruct, bool) {
	s.lock.RLock()
	value, ok := s.m[key]
	s.lock.RUnlock()
	return value, ok
}

func (s *structCacheMap) Set(key reflect.Type, value *cachedStruct) {
	s.lock.Lock()
	s.m[key] = value
	s.lock.Unlock()
}
