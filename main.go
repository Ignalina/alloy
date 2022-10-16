package main

import (
	"fmt"
	"github.com/apache/arrow/go/v9/arrow"
	"github.com/apache/arrow/go/v9/arrow/array"
	"github.com/apache/arrow/go/v9/arrow/memory"
)

func main() {
	mem := memory.NewGoAllocator()

	bld0 := array.NewInt32Builder(mem)
	defer bld0.Release()
	bld0.AppendValues([]int32{122}, []bool{true})
	arr0 := bld0.NewInt32Array() // materialize the array
	defer arr0.Release()

	bld1 := array.NewInt64Builder(mem)
	defer bld1.Release()
	bld1.AppendValues([]int64{122}, []bool{true})
	arr1 := bld1.NewInt64Array() // materialize the array
	defer arr1.Release()

	listOfarrays := []arrow.Array{arr1, arr1}

	fmt.Printf("[Go]\tCalling the goBridge with:\n\tarr: %v\n", listOfarrays)

	goBridge := GoBridge{GoAllocator: mem}
	i, err := goBridge.From_chunks(listOfarrays)

	if nil != err {
		fmt.Println(err)
	} else {
		fmt.Printf("[Go]\tRust counted %v arrays sent through ffi\n", i)
	}
}
