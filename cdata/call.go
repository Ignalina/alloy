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

func (goBridge GoBridge) Call(arrays []*array.Int32) (int, error) {
    var carrowschemas []cdata.CArrowSchema
    var carrowarrays []cdata.CArrowArray
    for idx, array := range arrays {
        fmt.Printf("[Go]\tExporting schema+array #%v\n", idx + 1)
        cas := cdata.CArrowSchema{}
        caa := cdata.CArrowArray{}
        cdata.ExportArrowArray(array, &caa, &cas) 
        carrowschemas = append(carrowschemas, cas)
        carrowarrays = append(carrowarrays, caa)
    }

	fmt.Printf("[Go]\tCalling Rust through C ffi now...\n")
    ret := C.call_with_ffi_voidptr(
        unsafe.Pointer(&carrowschemas),
        unsafe.Pointer(&carrowarrays),
        C.uintptr_t(len(carrowschemas)),
    )

	fmt.Printf("[Go]\tHello, again! Successfully sent Arrow data to Rust.\n")
	return int(ret), nil
}

