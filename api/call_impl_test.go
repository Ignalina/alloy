package api

import (
	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/memory"
	"testing"
)

func TestHelloName(t *testing.T) {

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "f1-i32", Type: arrow.PrimitiveTypes.Int32},
			{Name: "f2-f64", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	)

	goBridge := GoBridge{GoAllocator: memory.NewGoAllocator()}
	err := goBridge.CallWithTable(schema)

	if err != nil {
		t.Fatalf(`CallWithTable(nil) = %q `, err)
	}
}