package api

import (
	"fmt"
)

/*
#cgo LDFLAGS: -L../lib/libimpl.a -ldl
#include "../lib/impl.h"
#include "ARROW_C_DATA_INTERFACE.h"
*/
import "C"

/*

The callwithtable func wants the following type 
as argument:
    *_Ctype_struct___0

Is the arrow Schema that type? probably not.,
But we can create a reference to the C typedef
struct ArrowSchema and pass that to our function
and everyone is happy!
I dont know how to test this though? only how to 
build it haha
*/
func main() {
    arrowschema := &C.ArrowSchema{}
    C.callwithschema(arrowschema)
    fmt.Printf("Hello from Go!")
}

