package api

import (
	"fmt"
	"github.com/apache/arrow/go/v9/arrow"
	"github.com/apache/arrow/go/v9/arrow/cdata"
	"github.com/apache/arrow/go/v9/arrow/memory"
	"time"
	"unsafe"
)

/*
#cgo LDFLAGS: ./ffi/librust_impl.a -ldl -lm
#include "../cdata/arrow/c/abi.h"
#include "../ffi/impl.h"
*/
import "C"

func Info(s string) {
	t := time.Now()
	fmt.Printf(
		"[%d-%d-%d %d:%d:%d] [INFO] [Go]\t%s\n",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		s,
	)
}

type GoBridge struct {
	GoAllocator *memory.GoAllocator
}

func (goBridge GoBridge) FromChunks(arrays []arrow.Array) (int, error) {
	var Cschemas []cdata.CArrowSchema
	var Carrays []cdata.CArrowArray

	for idx, array := range arrays {
		Info(fmt.Sprintf("Exporting ArrowSchema and ArrowArray #%d to C", idx+1))
		cas := cdata.CArrowSchema{}
		caa := cdata.CArrowArray{}
		cdata.ExportArrowArray(array, &caa, &cas)
		Cschemas = append(Cschemas, cas)
		Carrays = append(Carrays, caa)
	}

	Info(fmt.Sprintf("Calling Rust through C ffi now with %v ArrowArrays", len(Cschemas)))
	ret := C.call_with_ffi_voidptr(
		unsafe.Pointer(&Cschemas[0]),
		unsafe.Pointer(&Carrays[0]),
		C.uintptr_t(len(Cschemas)),
	)

	Info("Hello, again! Successfully sent Arrow data to Rust.")
	return int(ret), nil
}
