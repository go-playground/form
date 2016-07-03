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
	timeType = reflect.TypeOf(time.Time{})
)
