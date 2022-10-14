use libc::c_int;
use arrow::ffi::{FFI_ArrowArray, FFI_ArrowSchema};
use arrow::array::{Array, ArrayRef, make_array_from_raw};

#[no_mangle]
pub extern "C" fn call_with_ffi(
    ffi_array: *const FFI_ArrowArray,
    ffi_schema: *const FFI_ArrowSchema
 ) -> c_int {

    let _arr = unsafe { make_array_from_raw(ffi_array, ffi_schema); };
    1 as c_int
}

#[no_mangle]
pub extern "C" fn call_with_ffi_schema(
    ffi_schema: *const FFI_ArrowSchema
) -> c_int {

    println!("Hello from Rust! Here is address to schema: {:?}", ffi_schema); 
    1 as c_int
}
