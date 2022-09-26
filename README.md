# Alloy
Go (Arrow buffs)--> Rust 
calls with Apache Arrow datatype's as parameter

Under early setup at atm reading these as inspiration.
*  https://github.com/mediremi/rust-plus-golang 
* https://michael-f-bryan.github.io/rust-ffi-guide/cbindgen.html to generate an extern C callable from GO
* https://arrow.apache.org/docs/status.html#ipc-format

You need Go/Rust and build tools installed. Try  
prompt> make build-all

it currently fails ..help me fix it..

Goals

## V0 Either using C-api or IPC
 
* Transfer hardcoded array of floats / ints/ strings..  
* Transfer schema describing array of arrays datatypes ... (Table!)  
* Buildscripts using Docker building GO and possible RUST part  

## V0.1   

##n Thanks to inspiring GO-Rust repo
* https://github.com/mediremi/rust-plus-golang
