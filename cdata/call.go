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
int from_chunks_ffi(const struct ArrowArray *arrptr, const struct  ArrowSchema *schptr, uintptr_t l);
int call_with_ffi_voidptr(void* schema, void* array, uintptr_t l) {
    return from_chunks_ffi(array, schema, l);
}

*/
import "C"

type GoBridge struct {
	GoAllocator *memory.GoAllocator
}

func (goBridge GoBridge) Call(array0 array.Int32, array1 array.Int64) (int, error) {
	fmt.Printf("[Go]\tCalling Rust through C ffi now...\n")
	var cas [2]cdata.CArrowSchema
	var caa [2]cdata.CArrowArray

	cdata.ExportArrowArray(&array0, &caa[0], &cas[0])
	cdata.ExportArrowArray(&array1, &caa[1], &cas[1])

	i := C.call_with_ffi_voidptr(unsafe.Pointer(&cas), unsafe.Pointer(&caa), C.uintptr_t(2))

	fmt.Printf("[Go]\tHello, again! Successfully sent Arrow data to Rust.\n")
	return int(i), nil
}
