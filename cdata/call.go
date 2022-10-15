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
void from_chunks_ffi(const struct ArrowArray *arrptr, const struct  ArrowSchema *schptr, uintptr_t l);

void call_with_ffi_voidptr(void* schema,void* array,uintptr_t l) {
from_chunks_ffi(array,schema,l);
}

*/
import "C"

type GoBridge struct {
	GoAllocator *memory.GoAllocator
}

func (goBridge GoBridge) Call(array array.Int32) error {
	fmt.Printf("Hello from Go! Calling Rust through C ffi now...\n")
	var cas [1]cdata.CArrowSchema
	var caa [1]cdata.CArrowArray

	//	cas := &cdata.CArrowSchema{}
	//	caa := &cdata.CArrowArray{}
	cdata.ExportArrowArray(&array, &caa[0], &cas[0])

	fmt.Printf("You can do it , go Rust land !\n")

	C.call_with_ffi_voidptr(unsafe.Pointer(&cas), unsafe.Pointer(&caa), C.uintptr_t(1))

	fmt.Printf("Hello0 from Go, again! Successfully sent Arrow data to Rust.\n")
	return nil
}
