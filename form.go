package form

import (
	"reflect"
	"time"
)

const (
	blank              = ""
	namespaceSeparator = '.'
	ignore             = "-"
	fieldNS            = "Field Namespace:"
	errorText          = " ERROR:"
)

var (
	// once     sync.Once
	// nsPool   *sync.Pool
	timeType = reflect.TypeOf(time.Time{})
)

// func init() {
// 	once.Do(func() {
// 		nsPool = &sync.Pool{New: func() interface{} {
// 			return make([]byte, 0, 64)
// 		}}
// 	})
// }
