package main

import (
	"fmt"
	"github.com/apache/arrow/go/v9/arrow"
	"github.com/apache/arrow/go/v9/arrow/array"
	"github.com/apache/arrow/go/v9/arrow/memory"
	arrow2 "github.com/ignalina/alloy/cdata"
)

func main() {
	mem := memory.NewGoAllocator()

	builder := array.NewInt32Builder(mem)
	builder.AppendValues([]int32{122}, nil)

	arr := builder.NewInt32Array()
	defer arr.Release()

	sch := arrow.NewSchema(
		[]arrow.Field{
			{Name: "f1-i32", Type: arrow.PrimitiveTypes.Int32},
			{Name: "f2-i32", Type: arrow.PrimitiveTypes.Int32}},
		nil,
	)

	fmt.Printf("Calling the goBridge with:\narray=%v\nschema=%v\n", arr, sch)
	goBridge := arrow2.GoBridge{GoAllocator: mem}
	err := goBridge.Call(arr, sch)
	fmt.Println(err)
}
