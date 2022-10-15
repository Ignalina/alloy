package main

import (
	"fmt"
	"github.com/apache/arrow/go/v9/arrow/array"
	"github.com/apache/arrow/go/v9/arrow/memory"
	"github.com/ignalina/alloy/cdata"
)

func main() {
	mem := memory.NewGoAllocator()

	bld := array.NewInt32Builder(mem)
	defer bld.Release()
	bld.AppendValues([]int32{122}, []bool{true})
	arr := bld.NewInt32Array() // materialize the array
	defer arr.Release()

	fmt.Printf("Calling the goBridge with:\narray=%v\n", arr)
	goBridge := cdata.GoBridge{GoAllocator: mem}
	err := goBridge.Call(*arr)

	if (nil != err) {
		fmt.Println(err)
	}
}
