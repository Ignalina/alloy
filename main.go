package main

import (
	"fmt"
	"github.com/apache/arrow/go/v9/arrow/array"
	"github.com/apache/arrow/go/v9/arrow/memory"
	"github.com/ignalina/alloy/cdata"
)

func appendToBuilder(builder *array.Int32Builder, values []int32) {
    var valids []bool
    for idx := 0; idx < len(values); idx++ {
        valids = append(valids, true)
    }
    builder.AppendValues(values, valids)
}

func main() {
	mem := memory.NewGoAllocator()
    values := []int32{1, 2, 3, -4}
    var arrays []*array.Int32

    // Append values to the Int32 Builder,
    // then materialize the Array in memory.
	bld0 := array.NewInt32Builder(mem)
    appendToBuilder(bld0, values)
	arr0 := bld0.NewInt32Array()
	defer bld0.Release()
	defer arr0.Release()

	bld1 := array.NewInt32Builder(mem)
    appendToBuilder(bld1, values)
	arr1 := bld1.NewInt32Array()
    defer bld1.Release()
	defer arr1.Release()

    arrays = append(arrays, arr0)
    arrays = append(arrays, arr1)
    fmt.Printf("[Go]\tCalling the goBridge with:\n\t%v\n", arrays)

	goBridge := cdata.GoBridge{GoAllocator: mem}
	ret, err := goBridge.Call(arrays)

	if nil != err {
		fmt.Println(err)
	} else {
		fmt.Printf("[Go]\tRust counted %v arrays sent through ffi\n", ret)
	}
}
