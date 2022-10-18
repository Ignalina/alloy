![Builds](https://github.com/Ignalina/alloy/actions/workflows/builds.yml/badge.svg)
![Rust tests](https://github.com/Ignalina/alloy/actions/workflows/rust-tests.yml/badge.svg)
![Go tests](https://github.com/Ignalina/alloy/actions/workflows/go-tests.yml/badge.svg)

![alloy](https://raw.githubusercontent.com/Ignalina/alloy/main/images/alloy.svg)<br>

An GO module allowing calls to Rust code code with Apache Arrow data. Alloy means a mixture between two or more components and is a joint venture between Wilhelm Ågren (Rust)  and Rickard Lundin(GO).

An example application in need of this is Thund , An Go/Rust based Apache Arrow centric DAG executor
The Apache Arrow data is either one-shot or IPC streaming

## Usage example
The example main.go should envision how your GO application utilize Alloy.

```golang
func main() {
    mem := memory.NewGoAllocator()
    values := [][]int32{
        {1, 2, 3, -4},
        {2, 3, 4, 5},
        {3, 4, 5, 6},
    }

    builders, arrays := buildAndAppend(mem, values)
    
	....
		
    goBridge := api.GoBridge{GoAllocator: mem}
    ret, err := goBridge.FromChunks(arrays)

    api.Info(fmt.Sprintf("Rust counted %v arrays sent through ffi", ret))
}

```

Output, where default dummy Rust backend echoes out the sent data. It will be up to the Rust guys to implent code to process the data to their hearts content !!


```bash
$ ./alloy
[2022-10-17 18:09:34] [INFO] [Go]	Exporting ArrowSchema and ArrowArray #1 to C
[2022-10-17 18:09:34] [INFO] [Go]	Exporting ArrowSchema and ArrowArray #2 to C
[2022-10-17 18:09:34] [INFO] [Go]	Exporting ArrowSchema and ArrowArray #3 to C
[2022-10-17 18:09:34] [INFO] [Go]	Calling Rust through C ffi now with 3 ArrowArrays
[2022-10-17 18:09:34] [INFO] [Rust]	Hello! Reading the ffi pointers now.
[2022-10-17 18:09:34] [INFO] [Rust]	array1: Int32[1, 2, 3, -4]
[2022-10-17 18:09:34] [INFO] [Rust]	array2: Int32[2, 3, 4, 5]
[2022-10-17 18:09:34] [INFO] [Rust]	array3: Int32[3, 4, 5, 6]
[2022-10-17 18:09:34] [INFO] [Go]	Hello, again! Successfully sent Arrow data to Rust.
[2022-10-17 18:09:34] [INFO] [Go]	Rust counted 3 arrays sent through ffi
```
## License
All code written is to be held under a general MIT-license, please see [LICENSE](https://github.com/Ignalina/alloy/blob/main/LICENSE) for specific information.


## Extended / Detailed readme below

![alloy](https://raw.githubusercontent.com/Ignalina/alloy/main/doc/alloy_schematic.svg)

## Goals and versions
In general, the overarching goal of the `alloy go module` is to enable Go to Rust calls through C
interface using Cgo and Rust ffi; with close to zero overhead using the Apache Arrow
data format. Only pointers referencing the allocated memory is sent between the
different language binaries, allowing for fast, (somewhat) robust, and colorful use
cases in data engineering scenarios.

### v0.1
- from Go to Rust Import Arrow Array chunks through ffi pointers to schema and array, .
- Send information back to Go instance from Rust.
- Access Go allocated memory without GC causing kernel panics.

### v0.2 (?)
- Aggregation on Arrow Array in Rust and accessing memory in Go.
- IPC (Streaming)

## Requirements
- Apache Arrow v9.0.0 https://arrow.apache.org/install/
- Go v1.19.1 https://go.dev/dl/ 
- Arrow2 v.0.14.2 https://crates.io/crates/arrow2

## Usage example 
Asume you have a MyCoolRustLib.a 
```golang
package main

import (
	"fmt"
	"github.com/apache/arrow/go/v9/arrow"
	"github.com/apache/arrow/go/v9/arrow/array"
	"github.com/apache/arrow/go/v9/arrow/memory"
	"github.com/ignalina/alloy"
)

// Assume you have 2 arrays arr1, arr2 ..

	listOfarrays := []arrow.Array{arr1, arr2}

	goBridge := GoBridge{GoAllocator: mem}
	goBridge.SetImplLib("/app/MyCoolRustLib.a","<BASE64 encoded cert>");
	
	i, err := goBridge.From_chunks(listOfarrays)

	if nil != err {
		fmt.Println(err)
	} else {
		fmt.Printf("[Go]\tRust counted %v arrays sent through ffi\n", i)
	}
```

## Setup (total rebuild)
If you are on a a debian based Linux system, you can very easily install the tools
needed to install Rust and similar tools, simply do this with the following command.
``` 
$ sudo apt-get install build-essential
```


You need the Go and Rust programming languages installed to even attempt at running
any of the code in this repository. If you are on a GNU/Linux based system, you can
run the comamnd `uname -ar` to see your specific OS version and CPU architecture. This
is important to know so that you download the correct Go version. Go (pun intended) to
this page https://go.dev/dl/ and download the correct file for your system. Next go
to the download location and run the following commands (checksum is optional, but a
sustainable practice),
```
$ sha256sum go1.19.1.linux-amd64.tar.gz
...
$ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.1.linux-amd64.tar.gz
...
$ export GOROOT=/usr/local/go
$ export GOPATH=$HOME/go
$ export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```
now you can verify that the Go compiler is installed properly by checking its version.
```
$ go version
``` 

Next step is to install Rust. Perhaps a bit counter-intuitive, but this is a lot
easier than the Go install. Simply go to this page https://rustup.rs/ and follow the
guide on screen. It will install the Rust toolchain which constitutes of; the Rust
compiler and the Cargo package manager. Or if you just want the command to run the
install, this is it.
```
$ curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```
Again you can verify that everything is installed properly by running
```
$ cargo version
```

and in theory, you could now make and run the examples in the repo. As for now, all
builds and tests are passing. 
```
$ make build-all
``` 

11 Seconds later on a 32 core threadipper PRO you will see
```bash
$ ./alloy
[2022-10-17 18:09:34] [INFO] [Go]	Exporting ArrowSchema and ArrowArray #1 to C
[2022-10-17 18:09:34] [INFO] [Go]	Exporting ArrowSchema and ArrowArray #2 to C
[2022-10-17 18:09:34] [INFO] [Go]	Exporting ArrowSchema and ArrowArray #3 to C
[2022-10-17 18:09:34] [INFO] [Go]	Calling Rust through C ffi now with 3 ArrowArrays
[2022-10-17 18:09:34] [INFO] [Rust]	Hello! Reading the ffi pointers now.
[2022-10-17 18:09:34] [INFO] [Rust]	array1: Int32[1, 2, 3, -4]
[2022-10-17 18:09:34] [INFO] [Rust]	array2: Int32[2, 3, 4, 5]
[2022-10-17 18:09:34] [INFO] [Rust]	array3: Int32[3, 4, 5, 6]
[2022-10-17 18:09:34] [INFO] [Go]	Hello, again! Successfully sent Arrow data to Rust.
[2022-10-17 18:09:34] [INFO] [Go]	Rust counted 3 arrays sent through ffi
```

## Reading these as inspiration and references:
- https://observablehq.com/@kylebarron/zero-copy-apache-arrow-with-webassembly
- https://github.com/mediremi/rust-plus-golang Rust code from Go using cgo and ffi
- https://michael-f-bryan.github.io/rust-ffi-guide/cbindgen.html to generate an extern C callable from GO
- https://arrow.apache.org/docs/status.html#ipc-format
- https://github.com/alexcrichton/rust-ffi-examples a lot of FFI examples, including go-to-rust
- https://stackoverflow.com/questions/23081990/using-empty-struct-properly-with-cgo some information regarding C structs in Go


