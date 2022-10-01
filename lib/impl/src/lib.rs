extern crate libc;
use arrow::ffi::FFI_ArrowSchema;

#[no_mangle]
pub extern "C" fn callwithschema(schema: &mut FFI_ArrowSchema)  {
    println!("Hello from Rust, with FFI_ArrowSchema: {:?}", schema);
}

