use libc::c_int;
use arrow::ffi::{FFI_ArrowArray, FFI_ArrowSchema};
use arrow::array::{Array, make_array_from_raw};
use arrow::datatypes::{DataType, Schema, Field};

#[no_mangle]
pub extern "C" fn call_with_ffi(
    ffi_array: &FFI_ArrowArray,
    ffi_schema: &FFI_ArrowSchema
 ) -> c_int {

    let _arr = unsafe { make_array_from_raw(ffi_array, ffi_schema); };
    1 as c_int
}

#[no_mangle]
pub extern "C" fn call_with_ffi_schema(
    ffi_schema: &FFI_ArrowSchema
) -> c_int {
    
    let schema: Schema = match Schema::try_from(ffi_schema) {
        Ok(s) => s,
        Err(e) => panic!("Could not convert to Rust schema: {:?}", e)
    };

    println!("Hello from Rust! Here is the schema: {:?}", schema); 
    1 as c_int
}
