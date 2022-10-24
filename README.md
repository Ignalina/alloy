![Builds](https://github.com/Ignalina/alloy/actions/workflows/builds.yml/badge.svg)
![Rust tests](https://github.com/Ignalina/alloy/actions/workflows/rust-tests.yml/badge.svg)
![Go tests](https://github.com/Ignalina/alloy/actions/workflows/go-tests.yml/badge.svg)

![alloy](https://raw.githubusercontent.com/Ignalina/alloy/main/images/alloy.svg)<br>

Full design/build/deploy doc se [doc/README](doc/README.md)

An GO module allowing calls to Rust code code with Apache Arrow data. Operate in either one-shot mode or IPC streaming.  
An example application in need of this is Thund , An Go/Rust based Apache Arrow centric DAG executor

Alloy means a mixture between two or more components and is a joint venture between Wilhelm Ågren (Rust)  and Rickard Lundin(GO).

## Usage example
The example main.go should envision how your GO application utilize Alloy.

```golang
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

	ret, err := b.FromChunks(arrays)

	if nil != err {
		fmt.Println(err)
	} else {
		rust.Info(fmt.Sprintf("Rust counted %v arrays sent through ffi", ret))
	}
}
```

Output, where default dummy Rust backend echoes out the sent data. It will be up to the Rust guys to implent code to process the data to their hearts content !!


```bash
$ ./alloy
[2022-10-17 18:09:34] [INFO] [Go]	Exporting ArrowSchema and ArrowArray #1 to C
[2022-10-17 18:09:34] [INFO] [Go]	Exporting ArrowSchema and ArrowArray #2 to C
[2022-10-17 18:09:34] [INFO] [Go]	Exporting ArrowSchema and ArrowArray #3 to C
[2022-10-17 18:09:34] [INFO] [Go]	Calling Rust through C rust now with 3 ArrowArrays
[2022-10-17 18:09:34] [INFO] [Rust]	Hello! Reading the rust pointers now.
[2022-10-17 18:09:34] [INFO] [Rust]	array1: Int32[1, 2, 3, -4]
[2022-10-17 18:09:34] [INFO] [Rust]	array2: Int32[2, 3, 4, 5]
[2022-10-17 18:09:34] [INFO] [Rust]	array3: Int32[3, 4, 5, 6]
[2022-10-17 18:09:34] [INFO] [Go]	Hello, again! Successfully sent Arrow data to Rust.
[2022-10-17 18:09:34] [INFO] [Go]	Rust counted 3 arrays sent through rust
```


## License
All code written is to be held under a general MIT-license, please see [LICENSE](https://github.com/Ignalina/alloy/blob/main/LICENSE) for specific information.
