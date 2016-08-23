package form

import (
	"reflect"
	"sync"
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
	m    atomic.Value // map[reflect.Type]*cachedStruct
	lock sync.Mutex
}

func newStructCacheMap() *structCacheMap {
	sc := new(structCacheMap)
	sc.m.Store(map[reflect.Type]*cachedStruct{})

	return sc
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

func (s *structCacheMap) parseStruct(mode Mode, current reflect.Value, key reflect.Type, tagName string) *cachedStruct {

	s.lock.Lock()

	// could have been multiple trying to access, but once first is done this ensures struct
	// isn't parsed again.
	cs, ok := s.Get(key)
	if ok {
		s.lock.Unlock()
		return cs
	}

	typ := current.Type()
	cs = &cachedStruct{fields: make([]cachedField, 0, 4)} // init 4, betting most structs decoding into have at aleast 4 fields.

	numFields := current.NumField()

	var fld reflect.StructField
	var name string

	for i := 0; i < numFields; i++ {

		fld = typ.Field(i)

		if fld.PkgPath != blank && !fld.Anonymous {
			continue
		}

		if name = fld.Tag.Get(tagName); name == ignore {
			continue
		}

		if mode == ModeExplicit && len(name) == 0 {
			continue
		}

		if len(name) == 0 {
			name = fld.Name
		}

		cs.fields = append(cs.fields, cachedField{idx: i, name: name})
	}

	s.Set(typ, cs)

	s.lock.Unlock()

	return cs
}
