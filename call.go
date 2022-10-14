package main

import (
    "fmt"
    "github.com/apache/arrow/go/v9/arrow"
    "github.com/apache/arrow/go/v9/arrow/array"
    "github.com/apache/arrow/go/v9/arrow/memory"
)

/*
#cgo LDFLAGS: ./ffi/librust_impl.a -ldl -lm
#include "./ffi/impl.h"
*/
import "C"

type GoBridge struct {
    GoAllocator *memory.GoAllocator
}

func exportArrowSchema(schema *arrow.Schema, out *C.struct_ArrowSchema) {
    field := schema.Fields()[0]
    out.dictionary = nil
    out.name = C.CString(field.Name)
    out.format = C.CString("i")
    out.metadata = (*C.char)(nil)
    out.flags = C.int64_t(0)
    out.n_children = C.int64_t(0)

    out.children = nil
}

func (goBridge GoBridge) Call(array *array.Int32, schema *arrow.Schema) error {
    fmt.Printf("Hello from Go! Calling Rust through C ffi now...\n")

    arrow_schema := &C.struct_ArrowSchema{}

    exportArrowSchema(schema, arrow_schema)

    ret := C.call_with_ffi_schema(arrow_schema)

    fmt.Printf("Hello from Go! Successfully called Rust with Arrow parameter. ret=%v.\n", ret)
    return nil
}
