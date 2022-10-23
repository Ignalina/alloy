/*
MIT License

Copyright (c) 2022 Wilhelm Ågren & Rickard Ernst Björn Lundin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package main

import (
	"fmt"
	"github.com/apache/arrow/go/v9/arrow"
	"github.com/apache/arrow/go/v9/arrow/array"
	memory "github.com/apache/arrow/go/v9/arrow/memory"
	"github.com/ignalina/alloy/api"
	"github.com/ignalina/alloy/ffi/rust"
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

	var b api.Bridge
	b = rust.Bridge{api.CommonParameter{mem}}

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

	ret, err := b.FromChunks(arrays)

	if nil != err {
		fmt.Println(err)
	} else {
		rust.Info(fmt.Sprintf("Rust counted %v arrays sent through ffi", ret))
	}
}
