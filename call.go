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

ArrowArray* dummy_arrow_array() {
    ArrowArray* arr = (ArrowArray*)malloc(sizeof(ArrowArray));

    return arr;
}

ArrowSchema* dummy_arrow_schema() {
    ArrowSchema* sch = (ArrowSchema*)malloc(sizeof(ArrowSchema));
    
    sch->format = "F1-i32";
    sch->name = "Testname";
    sch->metadata = "Testmetadata";

    return sch;

}
*/
import "C"

type GoBridge struct {
    GoAllocator *memory.GoAllocator
}

func (goBridge GoBridge) Call(array *array.Int32, schema *arrow.Schema) error {
    fmt.Printf("Hello from Go! Calling Rust through C ffi now...\n")

    arrow_array := C.dummy_arrow_array()
    arrow_schema := C.dummy_arrow_schema()

    fmt.Printf("ArrowArray=%v\n", arrow_array)
    fmt.Printf("ArrowSchema=%v\n", arrow_schema)
    fmt.Printf("=======================\n")

    /*
    arrow_schema := &C.ArrowSchema{
        C.CString("F1-i32"),            // format
        C.CString("testname"),          // name
        C.CString("metadata"),          // metadata
        C.int64_t(1),                      // flags
        C.int64_t(0),                      // n_children
        **C.ArrowSchema{&[0]byte{}},              // children
        *C.ArrowSchema{&[0]byte{}},               // dictionary
    }
    */
    C.call_with_ffi(arrow_array, arrow_schema)

    fmt.Printf("Hello from Go, again! Successfully sent Arrow data to Rust.\n")
    return nil
}
