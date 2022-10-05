package main

import (
	"fmt"
	"github.com/apache/arrow/go/v9/arrow"
	"github.com/apache/arrow/go/v9/arrow/array"
	"github.com/apache/arrow/go/v9/arrow/memory"
)

func main() {
    mem := memory.NewGoAllocator()

    builder := array.NewInt32Builder(mem)
    builder.AppendValues([]int32{122}, nil)

    arr := builder.NewInt32Array()
    defer arr.Release()

	schema := arrow.NewSchema(
		[]arrow.Field{{Name: "f1-i32", Type: arrow.PrimitiveTypes.Int32},},
		nil,
	)

	goBridge := GoBridge{GoAllocator: mem}

	err := goBridge.Call(arr, schema)
	fmt.Println(err)
}

