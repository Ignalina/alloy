package main

import (
	"fmt"
	"github.com/apache/arrow/go/v9/arrow"
	"github.com/apache/arrow/go/v9/arrow/array"
	"github.com/apache/arrow/go/v9/arrow/memory"
    "github.com/ignalina/alloy/api"
)

func appendToBuilder(builder *array.Int32Builder, values []int32) {
    var valids []bool
    for idx := 0; idx < len(values); idx++ {
        valids = append(valids, true)
    }
    builder.AppendValues(values, valids)
}

func buildAndAppend(mem *memory.GoAllocator, values [][]int32) ([]*array.Int32Builder, []arrow.Array) {
    var arrays []arrow.Array
    var builders []*array.Int32Builder
    num_vals := len(values)

    for idx := 0; idx < num_vals; idx++ {
        builder := array.NewInt32Builder(mem)
        appendToBuilder(builder, values[idx])
        array := builder.NewInt32Array()
        arrays = append(arrays, array)
        builders = append(builders, builder)
    }
    return builders, arrays
}

func main() {
	mem := memory.NewGoAllocator()
    values := [][]int32{
        {1, 2, 3, -4},
        {2, 3, 4, 5},
        {3, 4, 5, 6},
    }

    builders, arrays := buildAndAppend(mem, values)
    for idx := 0; idx < len(arrays); idx++ {
        defer builders[idx].Release()
        defer arrays[idx].Release()
    }

    fmt.Printf("[Go]\tCalling the goBridge with:\n\t%v\n", arrays)
	goBridge := api.GoBridge{GoAllocator: mem}
	ret, err := goBridge.FromChunks(arrays)

	if nil != err {
		fmt.Println(err)
	} else {
		fmt.Printf("[Go]\tRust counted %v arrays sent through ffi\n", ret)
	}
}

