package cdata

import (
	"fmt"
	"github.com/apache/arrow/go/v9/arrow"
	"github.com/apache/arrow/go/v9/arrow/array"
	"github.com/apache/arrow/go/v9/arrow/cdata"
	"github.com/apache/arrow/go/v9/arrow/memory"
	"unsafe"
)

/*
#cgo LDFLAGS: ./ffi/librust_impl.a -ldl -lm
#include "arrow/c/abi.h"
void ffi_call_schema(struct ArrowSchema* schema);
void ffi_call_schema_voidptr(void* schema) {
ffi_call_schema(schema);
}

*/
import "C"

type GoBridge struct {
	GoAllocator *memory.GoAllocator
}

func (goBridge GoBridge) Call(array *array.Int32, schema *arrow.Schema) error {
	fmt.Printf("Hello from Go! Calling Rust through C ffi now...\n")

	//arrow_array := C.dummy_arrow_array()
	//arrow_schema := C.dummy_arrow_schema()

	ll := &cdata.CArrowSchema{}
	cdata.ExportArrowSchema(schema, ll)

	fmt.Printf("=======================\n")

	// C.call_with_ffi(arrow_array, Csch)

	C.ffi_call_schema_voidptr(unsafe.Pointer(ll))

	fmt.Printf("Hello from Go, again! Successfully sent Arrow data to Rust.\n")
	return nil
}
