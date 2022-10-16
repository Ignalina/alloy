![alloy](https://github.com/Ignalina/alloy/blob/feature/readme/images/alloy.png)<br>
![Builds](https://github.com/Ignalina/alloy/actions/workflows/builds.yml/badge.svg)
![Rust tests](https://github.com/Ignalina/alloy/actions/workflows/rust-tests.yml/badge.svg)
![Go tests](https://github.com/Ignalina/alloy/actions/workflows/go-tests.yml/badge.svg)
---
Go (Arrow buffs) --> Rust calls with Apache Arrow datatype's as parameters.

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

## Setup
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
```
$ ./alloy
[Go]	Calling the goBridge with:
        array1: [122]
        array2: [122]
[Go]	Calling Rust through C ffi now...
[Rust]	Hello! Reading the ffi pointers now.
[Rust]	array1: Int32[122]
[Rust]	array2: Int64[122]
[Go]	Hello, again! Successfully sent Arrow data to Rust.
[Go]	Rust counted 2 arrays sent through ffi```
```

## License
All code written is to be held under a general MIT-license, please see 
[LICENSE](https://github.com/Ignalina/alloy/blob/main/LICENSE) for
specific information.

## Under early setup, reading these as inspiration and references:
- https://observablehq.com/@kylebarron/zero-copy-apache-arrow-with-webassembly
- https://github.com/mediremi/rust-plus-golang Rust code from Go using cgo and ffi
- https://michael-f-bryan.github.io/rust-ffi-guide/cbindgen.html to generate an extern C callable from GO
- https://arrow.apache.org/docs/status.html#ipc-format
- https://github.com/alexcrichton/rust-ffi-examples a lot of FFI examples, including go-to-rust
- https://stackoverflow.com/questions/23081990/using-empty-struct-properly-with-cgo some information regarding C structs in Go

