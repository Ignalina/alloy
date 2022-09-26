package api

import (
	"fmt"
	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/memory"
)

/*

#cgo LDFLAGS: ../lib/libimpl.a -ldl
#include "../lib/impl.h"


#include "ARROW_C_DATA_INTERFACE.c"

//extern int callwithtable(ArrowSchema* schema);

*/
import (
	"C"
)

type GoBridge struct {
	GoAllocator *memory.GoAllocator
}

// TODO add columns via arrays arrays
func (goBridge GoBridge) CallWithTable(schema *arrow.Schema) error {
	C.callwithtable(schema)
	fmt.Printf("hi")
	return nil
}
