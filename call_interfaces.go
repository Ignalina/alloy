package main

import (
	"github.com/apache/arrow/go/v9/arrow"
)

type call_with_ffi interface {
	Call(array *arrow.Array, schema *arrow.Schema) error
}

