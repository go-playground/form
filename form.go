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

// Mode specifies which mode the form decoder is to run
type Mode uint8

const (

	// ModeImplicit tries to parse values for all
	// fields that do not have an ignore '-' tag
	ModeImplicit Mode = iota

	// ModeExplicit only parses values for field with a field tag
	// and that tag is not the ignore '-' tag
	ModeExplicit
)
