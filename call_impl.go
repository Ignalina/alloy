package main

import (
	"fmt"
	"github.com/apache/arrow/go/v9/arrow"
	"github.com/apache/arrow/go/v9/arrow/array"
	"github.com/apache/arrow/go/v9/arrow/memory"
)

/*
#cgo LDFLAGS: ./lib/libimpl.a -ldl -lm
#include "./lib/impl.h"
#include "./ARROW_C_DATA_INTERFACE.h"
*/
import "C"

type GoBridge struct {
	GoAllocator *memory.GoAllocator
}

func (goBridge GoBridge) Call(arr *array.Int32, schema *arrow.Schema) error {
    fmt.Printf("Hello from Go! Calling Rust through C ffi now...")
    arrow_array := &C.ArrowArray{}
	arrow_schema := &C.ArrowSchema{}
	C.call_with_ffi(arrow_array, arrow_schema)
	fmt.Printf("Hello from Go, again! I am done now.")
	return nil
}

