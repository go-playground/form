package form

import (
	"reflect"
	"sync/atomic"
)

type cachedField struct {
	idx  int
	name string
}

type cachedStruct struct {
	fields []cachedField
}

type structCacheMap struct {
	m atomic.Value // map[reflect.Type]*cachedStruct
}

func (s *structCacheMap) Get(key reflect.Type) (value *cachedStruct, ok bool) {
	value, ok = s.m.Load().(map[reflect.Type]*cachedStruct)[key]
	return
}

func (s *structCacheMap) Set(key reflect.Type, value *cachedStruct) {

	m := s.m.Load().(map[reflect.Type]*cachedStruct)

	nm := make(map[reflect.Type]*cachedStruct, len(m)+1)
	for k, v := range m {
		nm[k] = v
	}
	nm[key] = value
	s.m.Store(nm)
}
