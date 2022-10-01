package api

import (
	"fmt"
    "github.com/apache/arrow/go/v9/arrow"
    "github.com/apache/arrow/go/v9/arrow/memory"
)

/*
#cgo LDFLAGS: -L../lib/libimpl.a -ldl
#include "../lib/impl.h"
#include "ARROW_C_DATA_INTERFACE.h"
*/
import "C"

type GoBridge struct {
    GoAllocator *memory.GoAllocator
}

/*

The callwithtable func wants the following type 
as argument:
    *_Ctype_struct___0

Is the arrow Schema that type? probably not.,
But we can create a reference to the C typedef
struct ArrowSchema and pass that to our function
and everyone is happy!
*/
func main() {
    arrowschema := &C.ArrowSchema{}
    C.callwithschema(arrowschema)
    fmt.Printf("Hello from Go!")
}

func (goBridge GoBridge) Call(schema* arrow.Schema) error {
    arrowschema := &C.ArrowSchema{}
    C.callwithschema(arrowschema)
    fmt.Printf("Hello from Go!")
    return nil
}

