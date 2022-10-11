package cdata

import (
	"fmt"
	"github.com/apache/arrow/go/v9/arrow/array"
	"github.com/apache/arrow/go/v9/arrow/cdata"
	"github.com/apache/arrow/go/v9/arrow/memory"
	"unsafe"
)

/*
#cgo LDFLAGS: ./ffi/librust_impl.a -ldl -lm
#include "arrow/c/abi.h"

void call_with_ffi(struct ArrowSchema* schema,struct ArrowArray* array);
void call_with_ffi_voidptr(void* schema,void* array) {
call_with_ffi(schema,array);
}

*/
import "C"

type GoBridge struct {
	GoAllocator *memory.GoAllocator
}

func (goBridge GoBridge) Call(array *array.Int32) error {
	fmt.Printf("Hello from Go! Calling Rust through C ffi now...\n")

	cas := &cdata.CArrowSchema{}
	caa := &cdata.CArrowArray{}
	cdata.ExportArrowArray(array, caa, cas)

	fmt.Printf("You can do it , go Rust land !\n")

	C.call_with_ffi_voidptr(unsafe.Pointer(cas), unsafe.Pointer(caa))

	fmt.Printf("Hello0 from Go, again! Successfully sent Arrow data to Rust.\n")
	return nil
}
