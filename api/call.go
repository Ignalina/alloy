package api

import (
	"fmt"
	"github.com/apache/arrow/go/v9/arrow"
	"github.com/apache/arrow/go/v9/arrow/cdata"
	"github.com/apache/arrow/go/v9/arrow/memory"
	"unsafe"
)

/*
#cgo LDFLAGS: ./ffi/librust_impl.a -ldl -lm
#include "../cdata/arrow/c/abi.h"
#include "../ffi/impl.h"
*/
import "C"

type GoBridge struct {
	GoAllocator *memory.GoAllocator
}

func (goBridge GoBridge) FromChunks(arrays []arrow.Array) (int, error) {
    var Cschemas []cdata.CArrowSchema
    var Carrays []cdata.CArrowArray

    for idx, array := range arrays {
        fmt.Printf("[Go]\tExporting schema+array #%v\n", idx + 1)
        cas := cdata.CArrowSchema{}
        caa := cdata.CArrowArray{}
        cdata.ExportArrowArray(array, &caa, &cas) 
        Cschemas = append(Cschemas, cas)
        Carrays = append(Carrays, caa)
    }

    fmt.Printf("[Go]\tCalling Rust through C ffi now...\n")
    ret := C.call_with_ffi_voidptr(
        unsafe.Pointer(&Cschemas[0]),
        unsafe.Pointer(&Carrays[0]),
        C.uintptr_t(len(Cschemas)),
    )
    
    fmt.Printf("[Go]\tHello, again! Successfully sent Arrow data to Rust.\n")
    return int(ret), nil
}

