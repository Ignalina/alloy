/*
MIT License

Copyright (c) 2022 Wilhelm Ågren & Rickard Ernst Björn Lundin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package rust

import (
	"fmt"
	"github.com/apache/arrow/go/v9/arrow"
	"github.com/apache/arrow/go/v9/arrow/cdata"
	"github.com/ignalina/alloy/api"
	"time"
	"unsafe"
)

/*
#cgo LDFLAGS: -L${SRCDIR} -lrust_impl -ldl -lm
#include "cdata/arrow/c/abi.h"
#include "impl.h"
unsigned call_with_ffi_voidptr(void* schema, void* array, uintptr_t l)
{ return from_chunks_ffi(array, schema, l); }
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

type Bridge struct {
	api.CommonParameter
}

func (b Bridge) FromChunks(arrays []arrow.Array) (int, error) {
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
