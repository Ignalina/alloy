package main

import (
	"fmt"
	"github.com/apache/arrow/go/v9/arrow"
	"github.com/apache/arrow/go/v9/arrow/cdata"
	"github.com/apache/arrow/go/v9/arrow/memory"
	"unsafe"
)

/*
#cgo LDFLAGS: ./ffi/librust_impl.a -ldl -lm
#include "cdata/arrow/c/abi.h"
int from_chunks_ffi(const struct ArrowArray *arrptr, const struct  ArrowSchema *schptr, uintptr_t l);
int from_chunks_ffi_voidptr(void* schema, void* array, uintptr_t l) {
    return from_chunks_ffi(array, schema, l);
}

*/
import "C"

type GoBridge struct {
	GoAllocator *memory.GoAllocator
}

func (goBridge GoBridge) From_chunks(array []arrow.Array) (int, error) {
	fmt.Printf("[Go]\tCalling Rust through C ffi now...\n")

	// Yes Horrificly  using overisized constant array until we can transfer a dynamic sized slice/list via FFI
	var cas [100]cdata.CArrowSchema
	var caa [100]cdata.CArrowArray

	for i, _ := range array {
		cdata.ExportArrowArray(array[i], &caa[i], &cas[i])
	}

	handledRows := C.from_chunks_ffi_voidptr(unsafe.Pointer(&cas), unsafe.Pointer(&caa), C.uintptr_t(len(array)))

	fmt.Printf("[Go]\tHello, again! Successfully sent Arrow data to Rust.\n")
	return int(handledRows), nil
}
