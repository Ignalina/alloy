package api

import (
	"github.com/apache/arrow/go/v9/arrow"
	"github.com/apache/arrow/go/v9/arrow/memory"
	"testing"
	"C"
)
/*
#cgo LDFLAGS: -L../lib/libimpl.a -ldl
#include "../lib/impl.h"
#include "ARROW_C_DATA_INTERFACE.h"
*/


func TestHelloName(t *testing.T) {

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "f1-i32", Type: arrow.PrimitiveTypes.Int32},
			{Name: "f2-f64", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	goBridge := GoBridge{GoAllocator: memory.NewGoAllocator()}
	err := goBridge.Call(schema)

	if err != nil {
		t.Fatalf(`CallWithTable(nil) = %q `, err)
	}
}
