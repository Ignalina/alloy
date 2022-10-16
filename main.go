package main

import (
	"fmt"
	"github.com/apache/arrow/go/v9/arrow/array"
	"github.com/apache/arrow/go/v9/arrow/memory"
	"github.com/ignalina/alloy/cdata"
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

    fmt.Printf("[Go]\tCalling the goBridge with:\n\tarray1: %v\n\tarray2: %v\n", arr0, arr1)

	goBridge := cdata.GoBridge{GoAllocator: mem}
	i, err := goBridge.Call(*arr0, *arr1)

	if nil != err {
		fmt.Println(err)
	} else {
		fmt.Printf("[Go]\tRust counted %v arrays sent through FFI\n", i)
	}
}
